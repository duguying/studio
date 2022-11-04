// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/dbmodels"
	"duguying/studio/modules/imgtools"
	"duguying/studio/modules/storage"
	"duguying/studio/service/models"
	"duguying/studio/utils"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogather/com"
	"github.com/sirupsen/logrus"
)

func PutFile(c *CustomContext) (interface{}, error) {
	key := c.Query("path")
	if key == "" {
		return nil, fmt.Errorf("query path is required")
	}
	store := g.Config.Get("upload", "store-path", "store")
	fpath := filepath.Join(store, key)
	dir := filepath.Dir(fpath)
	err := com.MkdirWithCreatePath(dir)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(fpath)
	if err != nil {
		return nil, fmt.Errorf("create file failed, " + err.Error())
	}
	defer f.Close()

	written, err := io.Copy(f, c.Request.Body)
	if err != nil {
		return nil, fmt.Errorf("copy file failed, " + err.Error())
	}

	ext := path.Ext(key)
	mimeType := mime.TypeByExtension(ext)
	md5 := com.FileMD5(fpath)

	_, err = db.SaveFile(g.Db, key, mimeType, uint64(written), md5, c.UserID(), dbmodels.FileTypeArchive)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"ok": true,
	}, nil
}

// UploadImage 表单上传图片
// @Router /upload/image [post]
// @Tags 上传
// @Description 表单上传图片
// @Param publish body []byte true "图片内容"
// @Success 200 {object} models.UploadResponse
func UploadImage(c *CustomContext) (interface{}, error) {
	l := c.Logger()
	fh, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	// 图像是否优化，开启则调用 imagemagick 转码
	optimizeOption, optimize := c.GetPostForm("optimize")
	if optimize && optimizeOption == "false" {
		optimize = false
	}
	maxWidth := int64(0)
	scaleWidth, exist := c.GetPostForm("scale_width")
	if exist {
		maxWidth, _ = strconv.ParseInt(scaleWidth, 10, 32)
	}

	// 图片存储子目录
	storeDir := time.Now().Format("2006/01")
	uploadType, ok := c.GetPostForm("upload_type")
	if ok {
		if uploadType == "avatar" {
			storeDir = "avatar"
		}
	}

	// 图片存储根目录
	root := "img"
	_, private := c.GetPostForm("private")
	if private {
		root = "private"
	}

	// 文件存储根目录
	store := g.Config.Get("upload", "store-path", "store")

	// 文件信息
	size := fh.Size
	filename := strings.ToLower(fh.Filename)
	ext := filepath.Ext(filename)

	// 生成随机文件名，拼接路径
	randomName := utils.GenUID()
	key := filepath.Join(root, storeDir, fmt.Sprintf("%s%s", randomName, ext))
	fpath := filepath.Join(store, key)
	dir := filepath.Dir(fpath)
	if !com.PathExist(dir) {
		_ = os.MkdirAll(dir, 0644)
	}

	// 创建临时文件
	tdir := filepath.Join(filepath.Join(os.TempDir(), utils.GenUUID()), utils.GenUID())
	_ = os.MkdirAll(tdir, 0644)
	tpath := filepath.Join(tdir, filename)
	tf, err := os.Create(tpath)
	if err != nil {
		return nil, err
	}

	// 存储到临时文件
	hf, err := fh.Open()
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(tf, hf)
	if err != nil {
		return nil, err
	}
	hf.Close()
	tf.Close()

	// 获取宽度
	width, _, _ := imgtools.GetImgSize(tpath)

	// 转码
	log.Println("ext:", ext, "optimize:", optimize)
	optimizeSize := g.Config.GetInt64("image-optimize", "size", 512)
	if (imgNeedConvert(ext) || size >= 1024*optimizeSize || (width > maxWidth && maxWidth > 0)) && optimize {
		log.Println("ext optimize:", ext, "--> .webp")

		if width <= maxWidth {
			maxWidth = width
		}
		fpath = strings.TrimSuffix(fpath, ext) + ".webp"
		key = strings.TrimSuffix(key, ext) + ".webp"
		ext = ".webp"
		size, err = imgtools.ConvertImgToWebp(tpath, fpath, maxWidth)
		if err != nil {
			return nil, fmt.Errorf("转码失败, err:" + err.Error())
		}

		_ = os.RemoveAll(tdir)
	} else {
		err = utils.Movefile(tpath, fpath)
		if err != nil {
			return nil, err
		}
	}

	mimeType := mime.TypeByExtension(ext)
	md5 := com.FileMD5(fpath)

	// 存储文件信息到数据库
	fileInfo, err := db.SaveFile(g.Db, key, mimeType, uint64(size), md5, c.UserID(), dbmodels.FileTypeImage)
	if err != nil {
		return nil, err
	}

	go extractImgMetaAndSave(l, fileInfo.ID, fpath)

	// 存储到云存储
	costore, err := storage.NewCos(g.LogEntry.WithContext(c), storage.DefaultCosType)
	if err != nil {
		return nil, err
	}
	localPath := getLocalPath(fileInfo.Path)
	remotePath := fileInfo.Path
	err = costore.PutFile(localPath, remotePath)
	if err != nil {
		return nil, err
	}

	// 返回文件路径
	url := utils.GetFileURL(key)
	return models.UploadResponse{
		Ok:   true,
		URL:  url,
		Name: randomName,
	}, nil
}

func extractImgMetaAndSave(l *logrus.Entry, fileID, path string) {
	meta, metas, err := imgtools.ExtractImgMeta(path)
	if err != nil {
		l.Printf("extract image meta data failed, err: %s, path: %s\n", err.Error(), path)
		return
	}
	_, err = db.AddImageMeta(g.Db, fileID, meta, metas)
	if err != nil {
		l.Printf("add image meta data into db failed, err: %s\n", err.Error())
		return
	}
}

func UploadFile(c *CustomContext) (interface{}, error) {
	fh, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}
	store := g.Config.Get("upload", "store-path", "store")
	domain := g.Config.Get("upload", "file-domain", "http://file.duguying.net")
	size := fh.Size
	pdir := c.PostForm("project")
	key := ""
	if pdir == "" {
		key = filepath.Join(time.Now().Format("2006/01"), fh.Filename)
	} else {
		key = filepath.Join(pdir, fh.Filename)
	}

	fpath := filepath.Join(store, key)
	dir := filepath.Dir(fpath)
	_ = os.MkdirAll(dir, 0644)

	f, err := os.Create(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	hf, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer hf.Close()

	_, err = io.Copy(f, hf)
	if err != nil {
		return nil, err
	}

	ext := path.Ext(key)
	mimeType := mime.TypeByExtension(ext)
	md5 := com.FileMD5(fpath)

	_, err = db.SaveFile(g.Db, key, mimeType, uint64(size), md5, c.UserID(), dbmodels.FileTypeArchive)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"ok":  true,
		"url": domain + strings.Replace(filepath.Join("/", key), `\`, `/`, -1),
	}, nil
}

// PageFile 列举文件
// @Router /admin/file/list [get]
// @Tags 上传
// @Description 列举文件
// @Param page query string true "页码"
// @Param size query int true "每页数"
// @Success 200 {object} models.CommonResponse
func PageFile(c *CustomContext) (interface{}, error) {
	l := c.Logger()
	pageStr := c.Query("page")
	page, err := strconv.ParseUint(pageStr, 10, 32)
	if err != nil {
		return nil, err
	}

	sizeStr := c.Query("size")
	size, err := strconv.ParseUint(sizeStr, 10, 32)
	if err != nil {
		return nil, err
	}

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}

	list, total, err := db.PageFile(g.Db, page, size, c.UserID())
	if err != nil {
		log.Println("page file failed, err:", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return nil, err
	}

	cos, err := storage.NewCos(l, storage.QcloudCosType)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	apiList := []*models.File{}
	for _, item := range list {
		apiItem := item.ToModel()
		url := utils.GetFileURL(apiItem.Path)
		count, err := db.FileCountArticleRef(g.Db, url)
		if err != nil {
			continue
		}
		apiItem.ArticleRefCount = int(count)
		apiItem.LocalExist = com.FileExist(getLocalPath(apiItem.Path))
		coverRefCnt, err := db.FileCountCoverRef(g.Db, item.ID)
		if err != nil {
			continue
		}
		apiItem.CoverRefCount = int(coverRefCnt)

		wg.Add(1)
		go func(fileItem *models.File) {
			defer wg.Add(-1)
			exist, err := cos.IsExist(fileItem.Path)
			if err != nil {
				return
			}
			fileItem.COS = exist
		}(apiItem)

		apiList = append(apiList, apiItem)
	}
	wg.Wait()

	return models.FileAdminListResponse{
		Ok:    true,
		List:  apiList,
		Total: int(total),
	}, nil
}

func imgNeedConvert(ext string) bool {
	notNeedConvertMap := map[string]bool{
		".png": true,
	}
	_, ok := notNeedConvertMap[ext]
	return !ok
}

// FileLs 列举文件
// @Router /admin/file/ls [get]
// @Tags 上传
// @Description 列举文件
// @Param prefix query int true "路径前缀"
// @Success 200 {object} models.CommonResponse
func FileLs(c *CustomContext) (interface{}, error) {
	prefix := c.Query("prefix")
	list, err := db.ListCurrentDir(g.Db, c.UserID(), prefix)
	if err != nil {
		return nil, err
	}
	return &models.FileLsResponse{
		Ok:   true,
		List: list,
	}, nil
}

// FileSyncToCos 文件同步到COS
// @Router /admin/file/sync_cos [get]
// @Tags 上传
// @Description 文件同步到COS
// @Param page query string true "页码"
// @Param size query int true "每页数"
// @Success 200 {object} models.CommonResponse
func FileSyncToCos(c *CustomContext) (interface{}, error) {
	req := models.FileSyncRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		return nil, err
	}

	file, err := db.GetFile(g.Db, req.FileID)
	if err != nil {
		return nil, err
	}

	store, err := storage.NewCos(g.LogEntry.WithContext(c), req.CosType)
	if err != nil {
		return nil, err
	}

	localPath := getLocalPath(file.Path)
	remotePath := file.Path
	err = store.PutFile(localPath, remotePath)
	if err != nil {
		return nil, err
	}

	return models.CommonResponse{
		Ok: true,
	}, nil
}

func getLocalPath(path string) string {
	store := g.Config.Get("upload", "store-path", "store")
	return filepath.Join(store, path)
}

// DeleteFile 删除文件
func DeleteFile(c *CustomContext) (interface{}, error) {
	l := c.Logger()
	req := &models.FileSyncRequest{}
	err := c.BindQuery(&req)
	if err != nil {
		return nil, err
	}

	file, err := db.GetFile(g.Db, req.FileID)
	if err != nil {
		return nil, err
	}

	cos, err := storage.NewCos(l, storage.QcloudCosType)
	if err != nil {
		return nil, err
	}
	err = cos.RemoveFile(file.Path)
	if err != nil {
		return nil, err
	}

	cnt, err := db.CheckFileRef(g.Db, file)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, fmt.Errorf("该文件存在 %d 处引用，不能删除", cnt)
	}

	_ = os.Remove(getLocalPath(file.Path))
	err = db.DeleteFile(g.Db, req.FileID)
	if err != nil {
		return nil, err
	}

	return models.CommonResponse{
		Ok: true,
	}, nil
}

// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
	"duguying/studio/modules/storage"
	"duguying/studio/service/models"
	"duguying/studio/utils"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogather/com"
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

	err = db.SaveFile(g.Db, key, mimeType, uint64(written), md5, c.UserID())
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
	fh, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	// 图像是否优化，开启则调用 imagemagick 转码
	_, optimize := c.GetPostForm("optimize")

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
	_ = os.MkdirAll(dir, 0644)

	// 转码与转储
	log.Println("ext:", ext, "optimize:", optimize)
	if (imgNeedConvert(ext) || size >= 1024*1024) && optimize {
		log.Println("ext optimize:", ext, "--> .webp")
		hf, err := fh.Open()
		if err != nil {
			return nil, err
		}
		defer hf.Close()

		tdir := filepath.Join(os.TempDir(), utils.GenUID())
		_ = os.MkdirAll(tdir, 0644)

		tpath := filepath.Join(tdir, filename)
		f, err := os.Create(tpath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		_, err = io.Copy(f, hf)
		if err != nil {
			return nil, err
		}

		fpath = strings.TrimSuffix(fpath, ext) + ".webp"
		key = strings.TrimSuffix(key, ext) + ".webp"
		ext = ".webp"
		size, err = ConvertImgToWebp(tpath, fpath)
		if err != nil {
			return nil, fmt.Errorf("转码失败, err:" + err.Error())
		}
		_ = os.RemoveAll(tdir)
	} else {
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
	}

	mimeType := mime.TypeByExtension(ext)
	md5 := com.FileMD5(fpath)

	// 存储文件信息到数据库
	err = db.SaveFile(g.Db, key, mimeType, uint64(size), md5, c.UserID())
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

	err = db.SaveFile(g.Db, key, mimeType, uint64(size), md5, c.UserID())
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

	apiList := []*models.File{}
	for _, item := range list {
		apiList = append(apiList, item.ToModel())
	}

	return models.FileAdminListResponse{
		Ok:    true,
		List:  apiList,
		Total: int(total),
	}, nil
}

// ConvertImgToWebp 图片转码到webp
func ConvertImgToWebp(inpath string, outpath string) (size int64, err error) {
	cmd := exec.Command("convert", inpath, outpath)
	err = cmd.Run()
	if err != nil {
		return 0, err
	}
	return getFileSize(outpath)
}

func getFileSize(path string) (size int64, err error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func imgNeedConvert(ext string) bool {
	notNeedConvertMap := map[string]bool{
		".png": true,
	}
	_, ok := notNeedConvertMap[ext]
	return !ok
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

	store, err := storage.NewCos(req.CosType)
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

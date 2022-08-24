// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
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

func PutFile(c *gin.Context) {
	key := c.Query("path")
	if key == "" {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "query path is required",
		})
		return
	}
	store := g.Config.Get("upload", "store-path", "store")
	fpath := filepath.Join(store, key)
	dir := filepath.Dir(fpath)
	err := com.MkdirWithCreatePath(dir)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	f, err := os.Create(fpath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "create file failed, " + err.Error(),
		})
		return
	}
	defer f.Close()

	written, err := io.Copy(f, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "copy file failed, " + err.Error(),
		})
		return
	}

	ext := path.Ext(key)
	mimeType := mime.TypeByExtension(ext)
	md5 := com.FileMD5(fpath)

	err = db.SaveFile(g.Db, key, mimeType, uint64(written), md5)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

// PutImage 上传粘贴图片
// @Router /put/image [post]
// @Tags 上传
// @Description 上传粘贴图片
// @Param publish body []byte true "图片内容"
// @Success 200 {object} models.CommonResponse
func PutImage(c *CustomContext) (interface{}, error) {
	store := g.Config.Get("upload", "store-path", "store")
	name := c.GetHeader("name")
	mimeType := c.GetHeader("mime")
	ext := filepath.Ext(name)

	randomName := utils.GenUID()
	key := filepath.Join("img", time.Now().Format("2006/01"), fmt.Sprintf("%s%s", randomName, ext))
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

	if mimeType != "" {
		ext := path.Ext(key)
		mimeType = mime.TypeByExtension(ext)
	}
	md5 := com.FileMD5(fpath)

	err = db.SaveFile(g.Db, key, mimeType, uint64(written), md5)
	if err != nil {
		return nil, err
	}

	url := getImgRefURL(key)
	return models.UploadResponse{
		Ok:   true,
		URL:  url,
		Name: randomName,
	}, nil
}

func getImgRefURL(key string) string {
	imgHost := g.Config.Get("store", "img-host-url", "https://image.duguying.net")
	key = strings.TrimPrefix(key, "img")
	return imgHost + key
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
	_, optimize := c.GetPostForm("optimize")

	store := g.Config.Get("upload", "store-path", "store")
	size := fh.Size
	ext := strings.ToLower(filepath.Ext(fh.Filename))

	randomName := utils.GenUID()
	key := filepath.Join("img", time.Now().Format("2006/01"), fmt.Sprintf("%s%s", randomName, ext))
	fpath := filepath.Join(store, key)
	dir := filepath.Dir(fpath)
	_ = os.MkdirAll(dir, 0644)

	if imgNeedConvert(ext) && optimize {
		hf, err := fh.Open()
		if err != nil {
			return nil, err
		}
		defer hf.Close()

		tdir := filepath.Join(os.TempDir(), utils.GenUID())
		_ = os.MkdirAll(tdir, 0644)

		tpath := filepath.Join(tdir, fh.Filename)
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
		ext = ".webp"
		err = ConvertImgToWebp(tpath, fpath)
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

	err = db.SaveFile(g.Db, key, mimeType, uint64(size), md5)
	if err != nil {
		return nil, err
	}

	url := getImgRefURL(key)
	return models.UploadResponse{
		Ok:   true,
		URL:  url,
		Name: randomName,
	}, nil
}

func UploadFile(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
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
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	defer f.Close()

	hf, err := fh.Open()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}
	defer hf.Close()

	_, err = io.Copy(f, hf)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	ext := path.Ext(key)
	mimeType := mime.TypeByExtension(ext)
	md5 := com.FileMD5(fpath)

	err = db.SaveFile(g.Db, key, mimeType, uint64(size), md5)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":  true,
		"url": domain + strings.Replace(filepath.Join("/", key), `\`, `/`, -1),
	})
}

func PageFile(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.ParseUint(pageStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	sizeStr := c.Query("size")
	size, err := strconv.ParseUint(sizeStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}

	list, total, err := db.PageFile(g.Db, page, size)
	if err != nil {
		log.Println("page file failed, err:", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"list":  list,
		"total": total,
	})
	return
}

// ConvertImgToWebp 图片转码到webp
func ConvertImgToWebp(inpath string, outpath string) (err error) {
	cmd := exec.Command("convert", inpath, outpath)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func imgNeedConvert(ext string) bool {
	notNeedConvertMap := map[string]bool{
		".png": true,
	}
	_, ok := notNeedConvertMap[ext]
	return !ok
}

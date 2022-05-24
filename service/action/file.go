// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package action

import (
	"duguying/studio/g"
	"duguying/studio/modules/db"
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

func UploadImage(c *gin.Context) {
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
	key := filepath.Join("img", time.Now().Format("2006/01"), fmt.Sprintf("%s.png", utils.GenUID()))

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

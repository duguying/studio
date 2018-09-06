// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/30.

package action

import (
	"duguying/studio/g"
	"duguying/studio/service/db"
	"github.com/gin-gonic/gin"
	"github.com/gogather/com"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
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

	err = db.SaveFile(key, mimeType, uint64(written))
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
	size := fh.Size
	key := filepath.Join(time.Now().Format("2006/01"), fh.Filename)
	fpath := filepath.Join(store, key)

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

	err = db.SaveFile(key, mimeType, uint64(size))
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

	list, total, err := db.PageFile(page, size)
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
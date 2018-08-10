// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/8/10.

package action

import (
	"archive/tar"
	"compress/gzip"
	"duguying/studio/g"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogather/com"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CheckToken(c *gin.Context) {
	token := g.Config.Get("deployer", "token", "")
	reqToken := c.GetHeader("Token")
	if token == reqToken {
		c.Next()
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"ok":  false,
			"err": "auth failed",
		})
		c.Abort()
	}
	return
}

func PackageUpload(c *gin.Context) {
	appName := c.GetHeader("name")
	appPath := g.Config.Get("deployer", fmt.Sprintf("%s-path", appName), "" /*"/root/sites/parsing-techniques"*/)
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	if com.FileExist(appPath) {
		os.Rename(appPath, fmt.Sprintf("%s.%s", appPath, time.Now().Format("20060102150405")))
	} else {
		if !strings.HasSuffix(fh.Filename, ".tar.gz") {
			c.JSON(http.StatusOK, gin.H{
				"ok":  false,
				"err": "invalid file type",
			})
			return
		}

		fpath := fmt.Sprintf("%s.%s", appPath, ".tar.gz")
		f, err := os.Create(fpath)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ok":  false,
				"err": err.Error(),
			})
			return
		}

		hFile, err := fh.Open()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ok":  false,
				"err": err.Error(),
			})
			return
		}
		defer hFile.Close()

		_, err = io.Copy(f, hFile)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ok":  false,
				"err": err.Error(),
			})
			return
		}

		f.Close()

		// unzip file
		err = unzip(fpath)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"ok":  false,
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"ok": true,
			})
		}

		return

	}
}

func unzip(filePath string) error {
	// check file type
	if !strings.HasSuffix(filePath, ".tar.gz") {
		return fmt.Errorf("invalid file type")
	}

	// get prefix path
	prefix := strings.TrimSuffix(filePath, ".tar.gz")

	// file read
	fr, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	// tar read
	tr := tar.NewReader(gr)

	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// 显示文件
		log.Println("prefix", prefix)
		log.Println(h.Name)

		// 打开文件
		fw, err := os.OpenFile(filepath.Join(prefix, strings.TrimPrefix(h.Name, "./")), os.O_CREATE|os.O_WRONLY, 0644 /*os.FileMode(h.Mode)*/)
		if err != nil {
			panic(err)
		}
		defer fw.Close()

		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			panic(err)
		}

	}

	log.Println("un tar.gz ok")
	return nil
}

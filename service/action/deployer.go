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
		log.Println("get form file failed,", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	if com.FileExist(appPath) {
		os.Rename(appPath, fmt.Sprintf("%s.%s", appPath, time.Now().Format("20060102150405")))
	}

	if !strings.HasSuffix(fh.Filename, ".tar.gz") {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": "invalid file type",
		})
		return
	}

	fpath := fmt.Sprintf("%s.%s", appPath, "tar.gz")
	f, err := os.Create(fpath)
	if err != nil {
		log.Println("create file failed,", err.Error())
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
		log.Println("copy file failed,", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"err": err.Error(),
		})
		return
	}

	f.Close()

	// unzip file
	err = untgz(fpath, strings.TrimSuffix(fpath, ".tar.gz"))
	if err != nil {
		log.Println("untgz failed,", err.Error())
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

func untgz(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := filepath.Join(dest, hdr.Name)
		file, err := createFile(filename, os.FileMode(hdr.Mode), hdr.FileInfo().IsDir())
		if err != nil {
			return err
		}
		if file != nil {
			io.Copy(file, tr)
		}
	}
	return nil
}

func createFile(name string, perm os.FileMode, isDir bool) (*os.File, error) {
	if isDir {
		err := os.MkdirAll(name, perm)
		if err != nil {
			return nil, err
		} else {
			return nil, nil
		}
	} else {
		err := os.MkdirAll(filepath.Dir(name), perm)
		if err != nil {
			return nil, err
		}
		return os.Create(name)
	}
}

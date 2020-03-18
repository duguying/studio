// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"strings"
)

func GenUUID() (string, error) {
	guuid := uuid.NewV4()
	return strings.Replace(guuid.String(), "-", "", -1), nil
}

func HmacSha1(content string, key string) string {
	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(content))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	out.Seek(0, 0)
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func StrContain(keyword string, vendor []string) bool {
	for _, item := range vendor {
		if keyword == item {
			return true
		}
	}
	return false
}

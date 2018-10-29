// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"strings"
)

func GenUUID() (string, error) {
	guuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
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

func TitleToUri(title string) (uri string) {
	var buffer bytes.Buffer
	for _, c := range title {
		char := fmt.Sprintf("%c", c)
		pys := pinyin.LazyConvert(char, nil)
		pystr := strings.Join(pys, "")
		if pystr != "" {
			buffer.WriteString(pystr + "-")
		} else {
			buffer.WriteString(char)
		}
	}
	uri = buffer.String()
	uri = strings.Replace(uri, "\n", "", -1)
	uri = strings.Replace(uri, "\r", "", -1)
	uri = strings.Replace(uri, "\t", " ", -1)
	uri = strings.Replace(uri, " ", "-", -1)
	uri = strings.TrimSuffix(uri, "-")
	return uri
}

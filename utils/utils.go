// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package utils

import (
	"strings"
	"github.com/satori/go.uuid"
	"crypto/sha1"
	"fmt"
	"crypto/hmac"
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
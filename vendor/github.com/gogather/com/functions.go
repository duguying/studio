package com

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

const (
	_VERSION = "0.1.0307"
)

func Version() string {
	return _VERSION
}

// 获取用户头像
func GetGravatar(email string) string {
	return "http://www.gravatar.com/avatar/" + Md5(strings.ToUpper(email))
}

// 切割关键词为html片段
func TagSplit(keywords string) string {
	if "" == keywords {
		return ""
	}

	content := ""
	tags := strings.Split(keywords, ",")
	for _, value := range tags {
		content = content + fmt.Sprintf(`<a class="tags" href="/tag/%s/1">%s</a>,`, value, value)
	}
	return content
}

// 四舍五入
func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

// http GET
func Get(reqUrl string) (string, error) {
	response, err := http.Get(reqUrl)
	if nil != err {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		response.Body.Close()
		return "", err
	}
	return string(body), nil
}

package iplocation

import (
	"github.com/gogather/com"
	"io/ioutil"
	"net/http"
)

const (
	_VERSION = "0.1.0307"
)

func Version() string {
	return _VERSION
}

type IpLocation struct {
	key string
}

func NewIpLocation(key string) *IpLocation {
	return &IpLocation{key: key}
}

func (this *IpLocation) get(reqUrl string) (string, error) {
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

func (this *IpLocation) Location(ip string) (map[string]interface{}, error) {
	url := "http://api.ipinfodb.com/v3/ip-city/?key=" + this.key + "&ip=" + ip + "&format=json"
	jsonString, err := this.get(url)

	if err != nil {
		return nil, err
	}

	data, err := com.JsonDecode(jsonString)
	return data.(map[string]interface{}), err
}

package utils

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/gogather/com"
	"github.com/gogather/iplocation"
)

var il *iplocation.IpLocation

func init() {
	key := beego.AppConfig.String("iplocation_key")
	il = iplocation.NewIpLocation(key)
}

func GetLocation(ip string) (string, error) {
	json, err := il.Location(ip)
	if json == nil {
		return "", errors.New("json is nil")
	}
	countryName := json["countryName"].(string)
	regionName := json["regionName"].(string)
	cityName := json["cityName"].(string)
	data := map[string]interface{}{
		"countryName": countryName,
		"regionName":  regionName,
		"cityName":    cityName,
	}
	str, err := com.JsonEncode(data)
	return str, err
}

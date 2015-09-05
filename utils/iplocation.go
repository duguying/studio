package utils

import (
	"github.com/astaxie/beego"
	"github.com/gogather/com"
	"github.com/gogather/iplocation"
)

var il *iplocation.IpLocation

func init() {
	key := beego.AppConfig.String("iplocation_key")
	il = iplocation.NewIpLocation(key)
}

func GetLocation(ip string) string {
	json, _ := il.Location(ip)
	countryName := json["countryName"].(string)
	regionName := json["regionName"].(string)
	cityName := json["cityName"].(string)
	data := map[string]interface{}{
		"countryName": countryName,
		"regionName":  regionName,
		"cityName":    cityName,
	}
	str, _ := com.JsonEncode(data)
	return str
}

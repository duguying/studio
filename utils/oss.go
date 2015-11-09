package utils

import (
	"github.com/astaxie/beego"
	"github.com/yanunon/oss-go-api/oss"
)

var ossHost string
var ossInternal bool
var ossId string
var ossKey string
var ossBucket string
var c *oss.Client

func ossInit() {
	ossInternal, _ = beego.AppConfig.Bool("oss_internal")
	if ossInternal {
		ossHost = beego.AppConfig.String("oss_put_host_internal")
	} else {
		ossHost = beego.AppConfig.String("oss_put_host")
	}

	ossId = beego.AppConfig.String("oss_id")
	ossKey = beego.AppConfig.String("oss_key")
	ossBucket = beego.AppConfig.String("oss_bucket")

	c = oss.NewClient(ossHost, ossId, ossKey, 10)
}

func OssStore(opath, fpath string) error {
	ossInit()
	err := c.PutObject(ossBucket+"/"+opath, fpath)
	return err
}

func OssDelete(opath string) error {
	ossInit()
	err := c.DeleteObject(ossBucket + "/" + opath)
	return err
}

func OssGetURL(opath string) string {
	return beego.AppConfig.String("oss_get_host") + "/" + opath
}

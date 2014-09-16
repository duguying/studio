package controllers

import (
	. "blog/models"
	"blog/utils"
	// "encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// xmlrpc
type XmlrpcController struct {
	beego.Controller
}

func (this *XmlrpcController) Get() {
	str := utils.ReadFile("views/rpcxml/rsd.xml")
	host := beego.AppConfig.String("host")
	result := fmt.Sprintf(str, host, host)
	this.Ctx.WriteString(result)
	this.ServeXml()
}

func (this *XmlrpcController) Post() {
	result := ""
	params := utils.Unserialize(this.Ctx.Input.RequestBody)
	// log.Println(string(this.Ctx.Input.RequestBody))
	// log.Println(params)
	methodName, _ := utils.GetMethodName(string(this.Ctx.Input.RequestBody))
	log.Println(methodName)
	if "blogger.getUsersBlogs" == methodName {
		result = getBlog(params)
	} else if "metaWeblog.newPost" == methodName {
		result = newPost(params)
	} else if "wp.newCategory" == methodName {
		result = newCata(params)
	} else if "mt.setPostCategories" == methodName {
		result = setCata(params)
	} else if "metaWeblog.newMediaObject" == methodName {
		result = newMedia(params)
	} else if "metaWeblog.editPost" == methodName {
		result = editPost(params)
	} else if "blogger.deletePost" == methodName {
		result = deletePost(params)
	}

	this.Ctx.WriteString(result)
	this.ServeXml()
}

/////////////////////////////////////////////////////////////////

func login(username string, password string) bool {
	user, err := FindUser(username)
	if err != nil {
		return false
	} else {
		passwd := utils.Md5(password + user.Salt)
		if passwd == user.Password {
			return true
		} else {
			return false
		}
	}
}

func getBlog(params interface{}) string {
	username := params.([]interface{})[1].(string)
	password := params.([]interface{})[2].(string)
	result := login(username, password)

	if result {
		host := beego.AppConfig.String("host")
		str := utils.ReadFile("views/rpcxml/response_login.xml")
		return fmt.Sprintf(str, host+"/", 1, "独孤影", host+"/xmlrpc")
	} else {
		return utils.ReadFile("views/rpcxml/response_login_failed.xml")
	}
}

// 新建文章
func newPost(params interface{}) string {
	username := params.([]interface{})[1].(string)
	password := params.([]interface{})[2].(string)
	result := login(username, password)

	if result {
		str := utils.ReadFile("views/rpcxml/response_new_post.xml")

		title := params.([]interface{})[3].(map[string]interface{})["title"].(string)
		content := params.([]interface{})[3].(map[string]interface{})["description"].(string)
		keywords := ""
		categories := params.([]interface{})[3].(map[string]interface{})["categories"]
		if categories != nil {
			cata := categories.([]interface{})
			for _, v := range cata {
				keywords = fmt.Sprintf(keywords+"%s,", v.(string))
			}
			keywords = strings.TrimSuffix(keywords, ",")
		}

		id, err := AddArticle(title, content, keywords, username)

		if err == nil {
			return fmt.Sprintf(str, id)
		} else {
			str := utils.ReadFile("views/rpcxml/response_failed.xml")
			return fmt.Sprintf(str, "文章发布失败! 注意标题不能重名")
		}

	} else {
		return utils.ReadFile("views/rpcxml/response_login_failed.xml")
	}
}

// 新建Catalog
func newCata(params interface{}) string {
	username := params.([]interface{})[1].(string)
	password := params.([]interface{})[2].(string)
	result := login(username, password)

	name := params.([]interface{})[3].(map[string]interface{})["name"]
	id, _ := NewTag(name.(string))

	if result {
		str := utils.ReadFile("views/rpcxml/response_new_catalog.xml")
		return fmt.Sprintf(str, id)
	} else {
		return utils.ReadFile("views/rpcxml/response_login_failed.xml")
	}
}

func setCata(params interface{}) string {
	username := params.([]interface{})[1].(string)
	password := params.([]interface{})[2].(string)
	result := login(username, password)

	if result {
		return utils.ReadFile("views/rpcxml/response_set_post_catalog.xml")
	} else {
		return utils.ReadFile("views/rpcxml/response_login_failed.xml")
	}
}

// 新建媒体文件
func newMedia(params interface{}) string {
	username := params.([]interface{})[1].(string)
	password := params.([]interface{})[2].(string)
	result := login(username, password)

	if result {
		name := params.([]interface{})[3].(map[string]interface{})["name"]
		filetype := params.([]interface{})[3].(map[string]interface{})["type"]
		bits := params.([]interface{})[3].(map[string]interface{})["bits"]

		err := utils.ParseMedia("static/upload/"+name.(string), bits.(string))

		if nil != err {
			str := utils.ReadFile("views/rpcxml/response_failed.xml")
			return fmt.Sprintf(str, "上传写入失败")
		}

		// 文件保存到OSS
		t := time.Now()
		ossFilename := fmt.Sprintf("%d/%d/%d/%s", t.Year(), t.Month(), t.Day(), name.(string))
		err = utils.OssStore(ossFilename, "static/upload/"+name.(string))

		if nil != err {
			str := utils.ReadFile("views/rpcxml/response_failed.xml")
			return fmt.Sprintf(str, "图片保存到OSS失败")
		} else {
			os.Remove("./static/upload/" + name.(string))
			id, err := AddFile(name.(string), ossFilename, "oss", filetype.(string))
			if nil != err {
				log.Println(err)
				str := utils.ReadFile("views/rpcxml/response_failed.xml")
				return fmt.Sprintf(str, "图片信息添加到数据库失败")
			}
			str := utils.ReadFile("views/rpcxml/response_new_media_object.xml")
			return fmt.Sprintf(str, id, name.(string), utils.OssGetURL(ossFilename), filetype.(string))
		}
	} else {
		return utils.ReadFile("views/rpcxml/response_login_failed.xml")
	}

}

// 编辑(更新)文章
func editPost(params interface{}) string {
	strId := params.([]interface{})[0].(string)
	username := params.([]interface{})[1].(string)
	password := params.([]interface{})[2].(string)
	result := login(username, password)

	if result {

		title := params.([]interface{})[3].(map[string]interface{})["title"].(string)
		content := params.([]interface{})[3].(map[string]interface{})["description"].(string)
		keywords := ""
		categories := params.([]interface{})[3].(map[string]interface{})["categories"]
		if categories != nil {
			cata := categories.([]interface{})
			for _, v := range cata {
				keywords = fmt.Sprintf(keywords+"%s,", v.(string))
			}
			keywords = strings.TrimSuffix(keywords, ",")
		}

		id, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			str := utils.ReadFile("views/rpcxml/response_failed.xml")
			return fmt.Sprintf(str, "非法文章ID")
		}

		var newArt Article

		newArt.Title = title
		newArt.Keywords = keywords
		newArt.Content = content

		err = UpdateArticle(id, "", newArt)

		if err == nil {
			return utils.ReadFile("views/rpcxml/response_edit_post.xml")
		} else {
			str := utils.ReadFile("views/rpcxml/response_failed.xml")
			return fmt.Sprintf(str, "文章发布失败! 注意标题不能重名")
		}
	} else {
		return utils.ReadFile("views/rpcxml/response_login_failed.xml")
	}
}

// 删除文章
func deletePost(params interface{}) string {
	strId := params.([]interface{})[1].(string)
	username := params.([]interface{})[2].(string)
	password := params.([]interface{})[3].(string)
	result := login(username, password)

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		str := utils.ReadFile("views/rpcxml/response_failed.xml")
		return fmt.Sprintf(str, "非法文章ID")
	}

	if result {
		_, err := DeleteArticle(id, "")
		if nil != err {
			str := utils.ReadFile("views/rpcxml/response_failed.xml")
			return fmt.Sprintf(str, "文章删除失败!")
		} else {
			return utils.ReadFile("views/rpcxml/response_delete_post.xml")
		}
	} else {
		return utils.ReadFile("views/rpcxml/response_login_failed.xml")
	}
}

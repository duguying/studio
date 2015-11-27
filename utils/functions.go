package utils

import (
	"fmt"
	"github.com/gogather/com"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// 检查用户名
func CheckUsername(username string) bool {
	if username[0] >= '0' && username[0] <= '9' {
		return false
	}

	for i := 0; i < len(username); i++ {
		if !(username[i] == '_' ||
			(username[i] >= '0' && username[i] <= '9') ||
			(username[i] >= 'a' && username[i] <= 'z') ||
			(username[i] >= 'A' && username[i] <= 'Z')) {
			return false
		}
	}

	return true
}

// 获取用户头像
func GetGravatar(email string) string {
	return "http://www.gravatar.com/avatar/" + com.Md5(strings.ToUpper(email))
}

// 读取文件
func ReadFileByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}

// 切割关键词为html片段
func TagSplit(keywords string) string {
	if "" == keywords {
		return ""
	}

	content := ""
	tags := strings.Split(keywords, ",")
	for _, value := range tags {
		// fmt.Printf("arr[%d]=%d \n", index, value)
		content = content + fmt.Sprintf(`<a class="tags" href="/tag/%s/1">%s</a>,`, value, value)
	}
	return content
}

func WriteFile(fullpath string, str string) error {
	data := []byte(str)
	return ioutil.WriteFile(fullpath, data, 0644)
}

func GetDate(dateStr string) string {
	date, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return dateStr
	} else {
		return date.Format("2006-01-02")
	}
}

func GetDateCN(dateStr string) string {
	date, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return dateStr
	} else {
		return date.Format("2006年01月02日")
	}
}

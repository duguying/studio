package utils

import (
	"fmt"
	"github.com/Unknwon/com"
	"testing"
)

func TestTitleToUri(t *testing.T) {
	hans := "中国人is chinese"
	py := TitleToUri(hans)
	fmt.Println(py)
}

func TestGenUUID(t *testing.T) {
	test := "/art/1+1=2"
	fmt.Println(com.UrlEncode(test))
}

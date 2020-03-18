package utils

import (
	"fmt"
	"github.com/unknwon/com"
	"testing"
)

func TestGenUUID(t *testing.T) {
	test := "/art/1+1=2"
	fmt.Println(com.UrlEncode(test))
}

package utils

import (
	"fmt"
	"testing"

	"github.com/russross/blackfriday/v2"
	"github.com/unknwon/com"
)

func TestGenUUID(t *testing.T) {
	test := "/art/1+1=2"
	fmt.Println(com.UrlEncode(test))
}

func TestParseMath(t *testing.T) {
	content := "asdfa$放一$串中文就移位了sdf$$123$$dfgdf$$skdfjhkds$$ sdfs$$"

	content = ParseMath(string(content))
	content = string(blackfriday.Run([]byte(content)))

	// out := ParseMath(content)
	fmt.Println(content)
}

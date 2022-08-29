package utils

import (
	"fmt"
	"testing"

	"github.com/unknwon/com"
)

func TestGenUUID(t *testing.T) {
	test := "/art/1+1=2"
	fmt.Println(com.UrlEncode(test))
}

func TestParseMath(t *testing.T) {
	content := "asdfa放一串中文就移位了sdf$$123$$dfgdf$$skdfjhkds$$ sdfs$$"
	out := ParseMath(content)
	fmt.Println(out)
}

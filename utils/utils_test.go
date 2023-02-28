package utils

import (
	"fmt"
	"testing"

	"github.com/gogather/blackfriday/v2"
	"github.com/unknwon/com"
)

func TestGenUUID(t *testing.T) {
	test := "/art/1+1=2"
	fmt.Println(com.UrlEncode(test))
}

func markdownFull(input []byte) []byte {
	// set up the HTML renderer
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags:      blackfriday.CommonHTMLFlags,
		Extensions: blackfriday.CommonExtensions | blackfriday.LaTeXMath,
	})
	options := blackfriday.Options{
		Extensions: blackfriday.CommonExtensions | blackfriday.LaTeXMath,
	}
	return blackfriday.Markdown(input, renderer, options)
}

func TestParseMath(t *testing.T) {
	content := `asdfa$放一$串中文就移位了sdf$$123$$dfgdf$$skdfjhkds$$ sdfs$$

test

$$
a=b+c
$$
	`

	content = string(markdownFull([]byte(content)))

	// out := ParseMath(content)
	fmt.Println(content)
}

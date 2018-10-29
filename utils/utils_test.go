package utils

import (
	"fmt"
	"testing"
)

func TestTitleToUri(t *testing.T) {
	hans := "中国人is chinese"
	py := TitleToUri(hans)
	fmt.Println(py)
}

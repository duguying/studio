package utils

import (
	"fmt"
	"testing"
)

func TestLunar(t *testing.T) {
	lunar := NewLunar("1990-08-16", false)
	fmt.Println("==>", lunar)

	fmt.Println(LunarToSolar(lunar))
}

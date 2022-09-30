// Package utils 包注释
package utils

import (
	"testing"
	"time"
)

func TestGenerateICS(t *testing.T) {
	GenerateICS(
		"uuid",
		time.Now(), time.Now(), time.Hour,
		"总结标题",
		"南山区大新路艺华花园",
		"这是一个生日聚会",
		"https://duguying.net",
		"糖糖",
	)
}

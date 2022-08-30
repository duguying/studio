// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/5/18.

package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/martinlindhe/base36"
	"github.com/microcosm-cc/bluemonday"
	uuid "github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenUUID() (string, error) {
	guuid := uuid.NewV4()
	return strings.Replace(guuid.String(), "-", "", -1), nil
}

func HmacSha1(content string, key string) string {
	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(content))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	out.Seek(0, 0)
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func StrContain(keyword string, vendor []string) bool {
	for _, item := range vendor {
		if keyword == item {
			return true
		}
	}
	return false
}

var (
	base36map = map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14,
		'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
		'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24,
		'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
		'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34,
		'Z': 35,
	}
	base36mix = []rune{
		'L', '9', 'M', 'U', '7', 'B', '2', 'H', 'S', '3',
		'O', 'R', 'I', 'G', '5', 'K', 'Q', '6', 'J', 'T',
		'0', 'Y', 'N', '8', 'F', 'P', 'E', 'A', '1', 'Z',
		'D', 'W', 'V', 'X', '4', 'C',
	}
)

// GenUID 生成随机短号
func GenUID() string {
	uidMin := base36.Decode("10000000")
	uidMax := base36.Decode("zzzzzzzz")
	uid := rand.Intn(int(uidMax-uidMin)) + int(uidMin)
	b36s := base36.Encode(uint64(uid))
	mb36b := bytes.Buffer{}
	for _, c := range b36s {
		idx := base36map[c]
		mb36b.WriteRune(base36mix[idx])
	}
	return strings.ToLower(mb36b.String())
}

// TrimHTML 剔除HTML标签
func TrimHTML(content string) string {
	p := bluemonday.StripTagsPolicy()
	return p.Sanitize(content)
}

var (
	inlineMathReg, _ = regexp.Compile(`\$([\d\D][^\$]+)\$`)
)

// ParseMath 解析数学公式标签
func ParseMath(content string) string {
	count := 0
	out := ""
	rd := strings.NewReader(content)
	lexer := NewLexer(rd)
	for {
		start, pos, tok := lexer.Lex()
		out = out + string([]rune(content)[start:pos])

		if tok == EOF {
			break
		}
		if tok == MATH {
			count++
			if count%2 == 1 {
				out = out + "${1}" //`<div v-katex:auto>`
			} else if count%2 == 0 {
				out = out + "${0}" //`</div>`
				out = strings.ReplaceAll(out, "${1}", `<div v-katex:auto>`)
				out = strings.ReplaceAll(out, "${0}", `</div>`)
			}
		}
	}

	out = strings.ReplaceAll(out, "${1}", "$$")

	// 处理行内 math
	matches := inlineMathReg.FindAllString(out, -1)
	for _, match := range matches {
		policy := bluemonday.StripTagsPolicy()
		strippedMatch := policy.Sanitize(match)
		if strippedMatch == match {
			matchTmp := "<span v-katex:auto>" + strings.TrimPrefix(match, "$")
			matchTmp = strings.TrimSuffix(matchTmp, "$") + "</span>"
			out = strings.ReplaceAll(out, match, matchTmp)
		}
	}

	return out
}

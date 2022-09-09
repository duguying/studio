package imgtools

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ConvertImgToWebp 图片转码到webp
func ConvertImgToWebp(inpath string, outpath string, scaleWidth int64) (size int64, err error) {
	args := []string{}
	if scaleWidth > 0 {
		args = append(args, "-resize", fmt.Sprintf("%dx", scaleWidth))
	}
	args = append(args, inpath, outpath)
	cmd := exec.Command("convert", args...)
	err = cmd.Run()
	if err != nil {
		return 0, err
	}
	return getFileSize(outpath)
}

func getFileSize(path string) (size int64, err error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// GetImgSize 获取图片尺寸
func GetImgSize(path string) (width, height int64, err error) {
	// identify -ping -format '%w %h' /Users/rainesli/Desktop/F335F72D-3E57-4DE2-AE4F-947103583079.heic
	cmd := exec.Command("identify", "-ping", "-format", "%w %h", path)
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	segs := strings.Split(string(output), " ")
	if len(segs) < 2 {
		return 0, 0, fmt.Errorf("invalid output")
	}
	width, err = strconv.ParseInt(segs[0], 10, 32)
	if err != nil {
		return 0, 0, err
	}
	height, err = strconv.ParseInt(segs[1], 10, 32)
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}

// ExtractImgMeta 获取图片meta信息
func ExtractImgMeta(path string) (meta, metas string, err error) {
	cmd := exec.Command("convert", path, "json:")
	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}
	info := []interface{}{}
	err = json.Unmarshal(output, &info)
	if err != nil {
		return "", "", err
	}
	metas = string(output)

	if len(metas) > 0 {
		metaRaw, err := json.Marshal(metas[0])
		if err != nil {
			return "", "", err
		}
		meta = string(metaRaw)
	}

	return meta, metas, nil
}

// Package imgtools 图片处理工具库
package imgtools

import (
	"duguying/studio/utils"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gogather/com"
)

// ConvertImgToWebp 图片转码到webp
func ConvertImgToWebp(inpath string, outpath string, scaleWidth int64) (size int64, err error) {
	args := []string{"-limit", "memory", "100mb", "-limit", "map", "100mb"}
	if scaleWidth > 0 {
		args = append(args, "-resize", fmt.Sprintf("%dx", scaleWidth))
	}
	args = append(args, inpath, "-auto-orient", outpath)
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

	if len(info) > 0 {
		metaRaw, err := json.Marshal(info[0])
		if err != nil {
			return "", "", err
		}
		meta = string(metaRaw)
	}

	return meta, metas, nil
}

// MakeThumbnail 制作缩略图
func MakeThumbnail(path string, maxHeight int) (thumbKey string, err error) {
	if maxHeight <= 0 {
		return "", fmt.Errorf("invalid maxHeight")
	}

	args := []string{"-resize", fmt.Sprintf("x%d", maxHeight)}
	thumbKey = filepath.Join("img", "cache", fmt.Sprintf("%s.webp", utils.GenUUID()))
	thumbPath := utils.GetFileLocalPath(thumbKey)
	cacheDir := filepath.Dir(thumbPath)

	if !com.PathExist(cacheDir) {
		os.MkdirAll(cacheDir, 0644)
	}

	args = append(args, path, thumbPath)
	cmd := exec.Command("convert", args...)
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	return thumbKey, nil
}

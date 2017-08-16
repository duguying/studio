package com

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/toolkits/file"
)

// PathExist 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func PathExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// ReadFileByte 读取文件
func ReadFileByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}

// ReadFileString 读取文本文件
func ReadFileString(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd), err
}

// ReadFile 读取文本文件
func ReadFile(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	return string(fd), err
}

// FileExist 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// WriteFile 字符串写入文件
func WriteFile(fullpath string, str string) error {
	data := []byte(str)
	return ioutil.WriteFile(fullpath, data, 0644)
}

// Mkdir 创建文件夹
func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

// Copyfile
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func MkdirWithCreatePath(fullpath string) error {
	dirArray := strings.Split(fullpath, string(filepath.Separator))
	if len(dirArray) < 1 {
		return errors.New("fullpath mistake")
	} else if len(dirArray) == 1 {
		Mkdir(fullpath)
		return nil
	}

	path := ""
	var err error = nil
	for i := 0; i < len(dirArray); i++ {
		dir := dirArray[i]
		path = path + string(filepath.Separator) + dir
		p := SubString(path, 1, len(path)-1)
		if !PathExist(p) {
			err = Mkdir(p)
		}
	}

	return err
}

// 写入文件，若目录不存在则自动创建
func WriteFileWithCreatePath(fullpath string, str string) error {
	fileDir := ""
	var err error = nil
	if !FileExist(fullpath) {
		fileDir = filepath.Dir(fullpath)
		err = MkdirWithCreatePath(fileDir)
	}
	err = WriteFile(fullpath, str)

	return err
}

// 获取用户目录
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

// Dir 获取路径所在目录
func Dir(fullpath string) string {
	unixPath := strings.Replace(fullpath, fmt.Sprintf("%c", filepath.Separator), "/", -1)
	return strings.Replace(file.Dir(unixPath), "/", fmt.Sprintf("%c", filepath.Separator), -1)
}

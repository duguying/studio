package utils

import (
	"io"
	"net/http"
	"os"
)

func GetImage(urls string, path string) error {
	res, err := http.Get(urls)
	defer res.Body.Close()
	file, err := os.Create(path)
	io.Copy(file, res.Body)
	return err
}

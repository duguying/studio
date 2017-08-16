package com

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// json编码
func JsonEncode(data interface{}) (string, error) {
	a, err := json.Marshal(data)
	return string(a), err
}

// json解码
func JsonDecode(data string) (interface{}, error) {
	dataByte := []byte(data)
	var dat interface{}

	err := json.Unmarshal(dataByte, &dat)
	return dat, err
}

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

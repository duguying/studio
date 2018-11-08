package json

import "github.com/json-iterator/go"

func Marshal(v interface{}) (result []byte, err error) {
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	return jsons.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) (err error) {
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	return jsons.Unmarshal(data, v)
}

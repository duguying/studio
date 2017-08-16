package statistics

import (
	"errors"
	"github.com/gogather/com"
	"github.com/gogather/com/http"
	"github.com/gogather/com/log"
)

const (
	_VERSION = "0.1.0307"
)

func Version() string {
	return _VERSION
}

var langDataTotal map[string]interface{}
var httpClient *http.HTTPClient

func getHttpClien() *http.HTTPClient {
	if httpClient == nil {
		httpClient = http.NewHTTPClient()
	}

	return httpClient
}

func GetRepos(username string, token string) (string, error) {
	langDataTotal = make(map[string]interface{})

	jsonRepos, err := getHttpClien().Get("https://api.github.com/users/" + username + "/repos?access_token=" + token)

	if err != nil {
		log.Redln(err)
		return "", err
	}

	repos, err := com.JsonDecode(jsonRepos)

	if err != nil {
		log.Redln(err)
		return "", err
	}

	if msg, ok := repos.(map[string]interface{}); ok {
		return "", errors.New(msg["message"].(string))
	}

	for _, v := range repos.([]interface{}) {
		reposName := v.(map[string]interface{})["name"].(string)
		if err = getLangOfRepos(reposName, username, token); err != nil {
			return "", err
		}
	}

	ichartData := convert2Ichart(langDataTotal)
	json, err := com.JsonEncode(ichartData)

	if err != nil {
		return "", err
	}

	return json, nil
}

func getLangOfRepos(reposName string, username string, token string) error {
	jsonLangData, err := getHttpClien().Get("https://api.github.com/repos/" + username + "/" + reposName + "/languages?access_token=" + token)

	if err != nil {
		return err
	}

	langData, err := com.JsonDecode(jsonLangData)

	if err != nil {
		return err
	}

	if msg, ok := langData.(map[string]interface{}); ok {
		if _, ok = msg["message"]; ok {
			return errors.New(msg["message"].(string))
		}
	}

	for k, v := range langData.(map[string]interface{}) {
		if _, ok := langDataTotal[k]; ok {
			langDataTotal[k] = langDataTotal[k].(float64) + v.(float64)
		} else {
			langDataTotal[k] = v.(float64)
		}
	}

	return err
}

func convert2Ichart(data interface{}) interface{} {
	color := [15]string{"#FF8484", "#FFFF00", "#00FF00", "#00FFFF", "#0084FF", "#840000", "#8484C6", "#FF84FF", "#B44AB3", "#639B7F", "#C8B53B", "#BB5B47", "#C01CE1", "#A8F011", "#FFBA35"}
	length := len(data.(map[string]interface{}))
	objSlice := make([]interface{}, length)
	index := 0
	for k, v := range data.(map[string]interface{}) {
		objSlice[index] = interface{}(map[string]interface{}{"name": k, "value": v, "color": color[index%15]})
		index++
	}
	return objSlice
}

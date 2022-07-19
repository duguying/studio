// Package logger 日志适配器
package logger

import (
	"context"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type LogEntity interface {
	ID() string
	MustToJSON() string
}

func Report(entity LogEntity) {
	// 初始化 ES
	esConf := elasticsearch.Config{
		Addresses: []string{"http://jump.duguying.net:19200"},
	}
	es, err := elasticsearch.NewClient(esConf)
	if err != nil {
		panic(err)
	}
	es.Info()

	req := esapi.IndexRequest{
		Index:      "test",
		DocumentID: entity.ID(),
		Body:       strings.NewReader(entity.MustToJSON()),
		Refresh:    "true",
	}

	resp, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("ESIndexRequestErr: %s", err.Error())
		return
	}

	defer resp.Body.Close()
	if resp.IsError() {
		log.Printf("ESIndexRequestErr: %s", resp.String())
	} else {
		log.Printf("ESIndexRequestOk: %s", resp.String())
	}

}

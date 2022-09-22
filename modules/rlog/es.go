// Package rlog 日志适配器
package rlog

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type EsAdaptor struct {
	client *elasticsearch.Client
	addr   string
	index  string
}

func NewEsAdaptor(addr string, index string) (adaptor *EsAdaptor, err error) {
	ea := &EsAdaptor{
		addr:  addr,
		index: index,
	}
	err = ea.init()
	if err != nil {
		return nil, err
	}
	return ea, nil
}

func (ea *EsAdaptor) init() error {
	// 初始化 ES
	esConf := elasticsearch.Config{
		Addresses: []string{ea.addr},
	}
	es, err := elasticsearch.NewClient(esConf)
	if err != nil {
		return err
	}
	es.Info()
	ea.client = es
	return nil
}

func (ea *EsAdaptor) Close() {

}

func (ea *EsAdaptor) Report(line string) error {
	req := esapi.IndexRequest{
		Index:   ea.index,
		Body:    strings.NewReader(line),
		Refresh: "true",
	}

	resp, err := req.Do(context.Background(), ea.client)
	if err != nil {
		// log.Printf("ESIndexRequestErr: %s", err.Error())
		return err
	}

	defer resp.Body.Close()
	if resp.IsError() {
		return err
		// log.Printf("ESIndexRequestErr: %s", resp.String())
	} else {
		return nil
		// log.Printf("ESIndexRequestOk: %s", resp.String())
	}
}

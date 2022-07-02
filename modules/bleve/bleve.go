// Package bleve description
package bleve

import (
	"duguying/studio/g"
	"duguying/studio/modules/cron"
	"log"

	"github.com/blevesearch/bleve"
	"github.com/gogather/com"
)

// IndexInstance 打开索引实例
func IndexInstance(path string) (index bleve.Index, err error) {
	exist := com.PathExist(path)
	if exist {
		index, err = bleve.Open(path)
		if err != nil {
			return nil, err
		}
	} else {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(path, mapping)
		if err != nil {
			return nil, err
		}
	}
	return index, nil
}

func Init() {
	var err error
	g.Index, err = IndexInstance("bleve/article")
	if err != nil {
		log.Fatalln("open bleve index failed, err:", err.Error())
		return
	}

	cron.FlushArticleBleve()
}

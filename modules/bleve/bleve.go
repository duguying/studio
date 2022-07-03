// Package bleve description
package bleve

import (
	"duguying/studio/g"
	"duguying/studio/modules/cron"
	"log"

	"github.com/blevesearch/bleve"
	_ "github.com/blevesearch/bleve/analysis/analyzer/custom"
	"github.com/gogather/com"
	_ "github.com/wangbin/jiebago/tokenizers"
)

// IndexInstance 打开索引实例
func IndexInstance(path string) (index bleve.Index, err error) {
	indexMapping := bleve.NewIndexMapping()
	err = indexMapping.AddCustomTokenizer("jieba",
		map[string]interface{}{
			"file": "bleve/jiebago/dict.txt",
			"type": "jieba",
		})
	if err != nil {
		log.Fatal(err)
	}

	err = indexMapping.AddCustomAnalyzer("jieba",
		map[string]interface{}{
			"type":      "custom",
			"tokenizer": "jieba",
			"token_filters": []string{
				"possessive_en",
				"to_lower",
				"stop_en",
			},
		})
	if err != nil {
		log.Fatal(err)
	}

	indexMapping.DefaultAnalyzer = "jieba"

	exist := com.PathExist(path)
	if exist {
		index, err = bleve.Open(path)
		if err != nil {
			return nil, err
		}
	} else {
		// mapping := bleve.NewIndexMapping()
		index, err = bleve.New(path, indexMapping)
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

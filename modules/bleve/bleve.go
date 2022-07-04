// Package bleve description
package bleve

import (
	"duguying/studio/g"
	"duguying/studio/modules/cron"
	"log"

	"github.com/blevesearch/bleve"
	_ "github.com/blevesearch/bleve/analysis/analyzer/custom"
	"github.com/gogather/com"
	_ "github.com/ttys3/gojieba-bleve"
	"github.com/yanyiwu/gojieba"
)

// IndexInstance 打开索引实例
func IndexInstance(path string) (index bleve.Index, err error) {
	indexMapping := bleve.NewIndexMapping()
	err = indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     gojieba.DICT_PATH,
			"hmmpath":      gojieba.HMM_PATH,
			"userdictpath": gojieba.USER_DICT_PATH,
			"idf":          gojieba.IDF_PATH,
			"stop_words":   gojieba.STOP_WORDS_PATH,
			"type":         "gojieba",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	indexMapping.DefaultAnalyzer = "gojieba"

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

// Package bleve description
package bleve

import (
	"duguying/studio/g"
	"duguying/studio/modules/cron"
	"log"
	"path/filepath"

	"github.com/blevesearch/bleve"
	_ "github.com/blevesearch/bleve/analysis/analyzer/custom"
	"github.com/gogather/com"
	_ "github.com/ttys3/gojieba-bleve"
)

// IndexInstance 打开索引实例
func IndexInstance(path string) (index bleve.Index, err error) {
	dictDir := g.Config.Get("bleve", "gojieba-dict", "bleve/gojieba")
	dictPath := filepath.Join(dictDir, "jieba.dict.utf8")
	hmmPath := filepath.Join(dictDir, "hmm_model.utf8")
	userDictPath := filepath.Join(dictDir, "user.dict.utf8")
	idfPath := filepath.Join(dictDir, "idf.utf8")
	stopWordsPath := filepath.Join(dictDir, "stop_words.utf8")
	indexMapping := bleve.NewIndexMapping()
	err = indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     dictPath,
			"hmmpath":      hmmPath,
			"userdictpath": userDictPath,
			"idf":          idfPath,
			"stop_words":   stopWordsPath,
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

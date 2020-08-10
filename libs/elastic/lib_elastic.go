package libs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gin_frame/config"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

//Elas elasticsearh 客户端连接
var Elas *es7.Client

func init() {
	fmt.Print(config.LoadConfig().Elastic.Addr)
	cfg := elasticsearch.Config{
		Addresses: config.LoadConfig().Elastic.Addr,
	}
	var err error
	Elas, err = es7.NewClient(cfg)
	if err != nil {
		fmt.Println("elasticsearch err:", err)
	}
}

//CreateDocument 创建文档
func CreateDocument(index string, content string) {
	req := esapi.IndexRequest{
		Index: index,
		// DocumentID: documentID,
		Body:    strings.NewReader(content),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), Elas)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Print(res)
	defer res.Body.Close()
}

//Search 检索
func Search(query interface{}, index string) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := Elas.Search(
		Elas.Search.WithContext(context.Background()),
		Elas.Search.WithIndex(index),
		Elas.Search.WithBody(&buf),
		Elas.Search.WithTrackTotalHits(true),
		Elas.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Print(res)
	defer res.Body.Close()
}

package esclient

import (
	"github.com/olivere/elastic/v7"
	"log"
	"time"
)

var client *elastic.Client

func init() {
	var err error

	client, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5))

	if err != nil {
		log.Printf("현재 연결된 엘라스틱서치 노드가 없습니다.")
		panic(err)
	}
}

func Client() *elastic.Client {
	return client
}

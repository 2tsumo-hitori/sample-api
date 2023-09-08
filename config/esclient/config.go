package esclient

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

var client *elastic.Client

func InitElasticSearch() {
	var err error

	client, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5))

	if err != nil {
		panic(err)
	}

	fmt.Println(client)
}

func Client() *elastic.Client {
	return client
}

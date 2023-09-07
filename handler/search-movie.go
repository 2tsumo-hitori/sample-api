package handler

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"sample-api/config/esclient"
)

func SearchByKeyword[T any](searchKeyword string, resp *[]T) {
	query := elastic.NewMatchQuery("movieNm_text", searchKeyword)

	sendRequestToElastic(query, resp)
}

func AutoCompleteByKeyword[T any](searchKeyword string, resp *[]T) {
	query := elastic.
		NewBoolQuery().
		Should(elastic.
			NewMatchQuery("movieNm_ac", searchKeyword))

	sendRequestToElastic(query, resp)
}

func sendRequestToElastic[T any](query elastic.Query, resp *[]T) {
	client := esclient.Client()

	searchResult, err := client.
		Search().
		Index("my_movie_search").
		Query(query).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		log.Fatal(err)
	} else {
		for _, hit := range searchResult.Hits.Hits {
			var movie T

			if err := json.Unmarshal(hit.Source, &movie); err != nil {
				log.Fatal(err)
			}

			*resp = append(*resp, movie)
		}
	}
}

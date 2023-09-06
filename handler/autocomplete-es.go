package handler

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"sample-api/config/esclient"
	"sample-api/model"
)

func AutoComplete(searchKeyword string, resp *[]model.AutoCompleteResponse) {
	query := elastic.NewMatchQuery("movieNm_text", searchKeyword)

	client := esclient.Client()

	searchResult, err := client.Search().
		Index("my_movie_search").
		Query(query).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		log.Fatal(err)
	} else {
		for _, hit := range searchResult.Hits.Hits {
			movie := model.AutoCompleteResponse{}

			if err := json.Unmarshal(hit.Source, &movie); err != nil {
				log.Fatal(err)
			}

			*resp = append(*resp, movie)
		}
	}
}

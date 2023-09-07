package handler

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"sample-api/config/esclient"
	"sample-api/model"
	"sample-api/util"
)

const (
	index        = "my_movie_search"
	movieNmText  = "movieNm_text"
	word         = "word"
	movieNmAc    = "movieNm_ac"
	chosungFront = "movieNm_chosung_front"
	chosungBack  = "movieNm_chosung_back"
	movieNmCount = "movieNmCount"
)

func SearchByKeyword[T model.Response](searchKeyword string, resp *[]T) {
	query := elastic.NewMatchQuery(movieNmText, searchKeyword)

	sendRequestToElastic(query, resp)
}

func AutoCompleteByKeyword[T model.Response](searchKeyword string, resp *[]T) {
	query := queryBuilderByKeyword(searchKeyword)

	sendRequestToElastic(query, resp)
}

func sendRequestToElastic[T model.Response](query elastic.Query, resp *[]T) {
	client := esclient.Client()

	searchResult, err := client.
		Search().
		Index(index).
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

func queryBuilderByKeyword(searchKeyword string) elastic.Query {
	var query elastic.Query

	if spell, s := util.InspectSpell(searchKeyword); spell == word {
		query = elastic.
			NewBoolQuery().
			Should(elastic.
				NewMatchQuery(movieNmAc, s))
	} else {
		query = elastic.
			NewBoolQuery().
			Must(elastic.NewBoolQuery().
				Should(elastic.NewMatchQuery(chosungFront, s),
					elastic.NewMatchQuery(chosungBack, s))).
			Should(elastic.NewRangeQuery(movieNmCount).Lte(len(s)))
	}

	return query
}

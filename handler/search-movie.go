package handler

import (
	"context"
	"encoding/json"
	"github.com/2tsumo-hitori/sample-api/config/esclient"
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/2tsumo-hitori/sample-api/util"
	"github.com/olivere/elastic/v7"
	"log"
)

const (
	index           = "my_movie_search"
	movieNmText     = "movieNm_text"
	movieNmEngToKor = "movieNm_eng2kor"
	movieNmKorToEng = "movieNm_kor2eng"
	word            = "word"
	movieNmAc       = "movieNm_ac"
	chosungFront    = "movieNm_chosung_front"
	chosungBack     = "movieNm_chosung_back"
	movieNmCount    = "movieNmCount"
)

func SearchByKeyword[T model.Response](searchKeyword string, resp *[]T) {
	q := util.Queue{}

	_, s := util.InspectSpell(searchKeyword)

	q.Enqueue(elastic.NewMatchQuery(movieNmText, s))
	q.Enqueue(elastic.NewMatchQuery(movieNmEngToKor, s))
	q.Enqueue(elastic.NewMatchQuery(movieNmKorToEng, s))

	sendRequestToElastic(q, resp)
}

func AutoCompleteByKeyword[T model.Response](searchKeyword string, resp *[]T) {
	q := util.Queue{}

	q.Enqueue(queryBuilderByKeyword(searchKeyword))

	sendRequestToElastic(q, resp)
}

// 검색어를 일반검색, 한/영 오타변환, 영/한 오타변환 쿼리들로 만들어 queue로 쌓고 재귀적으로 오타교정 구현
// ** 오타교정 쿼리 추가예정 **
func sendRequestToElastic[T model.Response](query util.Queue, resp *[]T) {
	if query.IsEmpty() {
		return
	}

	client := esclient.Client()

	searchResult, err := client.
		Search().
		Index(index).
		Query(query.Dequeue().(elastic.Query)).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	if searchResult.TotalHits() == 0 {
		sendRequestToElastic(query, resp)
	}

	for _, hit := range searchResult.Hits.Hits {
		var movie T

		if err := json.Unmarshal(hit.Source, &movie); err != nil {
			log.Fatal(err)
		}

		*resp = append(*resp, movie)
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

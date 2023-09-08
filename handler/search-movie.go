package handler

import (
	"context"
	"encoding/json"
	"fmt"
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
	keyword := BuildSuggestQuery(searchKeyword)

	fmt.Println(keyword)

	q := util.Queue{}

	_, s := util.InspectSpell(searchKeyword)

	q.Enqueue(elastic.NewMatchQuery(movieNmText, s))
	q.Enqueue(elastic.NewMatchQuery(movieNmEngToKor, s))
	q.Enqueue(elastic.NewMatchQuery(movieNmKorToEng, s))

	sendRequestToElastic(q, resp)
}

func AutoCompleteByKeyword[T model.Response](searchKeyword string, resp *[]T) {
	q := util.Queue{}

	q.Enqueue(queryBuildByKeyword(searchKeyword))

	sendRequestToElastic(q, resp)
}

// 검색어를 [일반검색, 한/영 오타변환, 영/한 오타변환] 쿼리들로 만들어 queue에 쌓고 재귀적으로 오타교정 구현
// ** 오타교정 쿼리 추가예정 **
func sendRequestToElastic[T model.Response](queryQueue util.Queue, resp *[]T) {
	// 종료 조건
	if queryQueue.IsEmpty() {
		return
	}

	client := esclient.Client()

	searchResult, err := client.
		Search().
		Index(index).
		Query(queryQueue.Dequeue().(elastic.Query)).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	// 결과가 없을 경우 재귀
	if searchResult.TotalHits() == 0 {
		sendRequestToElastic(queryQueue, resp)
	}

	for _, hit := range searchResult.Hits.Hits {
		var movie T

		if err := json.Unmarshal(hit.Source, &movie); err != nil {
			log.Fatal(err)
		}

		*resp = append(*resp, movie)
	}
}

func queryBuildByKeyword(searchKeyword string) elastic.Query {
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

func BuildSuggestQuery(searchKeyword string) string {
	termSuggester := elastic.NewTermSuggester("movie-suggestion").
		Field("movieNm_text.spell").
		StringDistance("jaro_winkler").
		Text(searchKeyword)

	query := elastic.NewSearchSource().
		Suggester(termSuggester).
		FetchSource(false).
		TrackScores(true).
		Timeout("100ms")

	client := esclient.Client()

	searchResult, _ := client.Search().
		Index(index).
		SearchSource(query).
		Do(context.Background())

	suggestions := searchResult.Suggest["movie-suggestion"]

	var resp string

	for _, i := range suggestions {
		for _, j := range i.Options {
			resp = j.Text
		}
	}

	return resp
}

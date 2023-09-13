package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/2tsumo-hitori/sample-api/util"
	"github.com/olivere/elastic/v7"
	"log"
)

const (
	index        = "my_movie_search"
	word         = "word"
	movieNmAc    = "movieNm_ac"
	chosungFront = "movieNm_chosung_front"
	chosungBack  = "movieNm_chosung_back"
	movieNmCount = "movieNmCount"
)

// SendRequestToElastic 검색어를 [일반검색, 한/영 오타변환, 영/한 오타변환] 쿼리들로 만들어 queue에 쌓고 재귀적으로 오타교정 구현
func (es *DefaultElasticsearchService) SendRequestToElastic(queryQueue *util.Queue, resp *[]model.SearchResponse) {
	// 종료 조건
	if queryQueue.IsEmpty() {
		return
	}

	searchResult, err := es.client.
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
		es.SendRequestToElastic(queryQueue, resp)
	}

	for _, hit := range searchResult.Hits.Hits {
		var movie model.SearchResponse

		if err := json.Unmarshal(hit.Source, &movie); err != nil {
			log.Fatal(err)
		}

		*resp = append(*resp, movie)
	}
}

func (es *DefaultElasticsearchService) QueryBuildByKeyword(searchKeyword string) interface{} {
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

func (es *DefaultElasticsearchService) BuildSuggestQuery(suggestKeyword *string, ch chan bool) {
	termSuggester := elastic.NewTermSuggester("movie-suggestion").
		Field("movieNm_text.spell").
		StringDistance("jaro_winkler").
		Text(*suggestKeyword)

	query := elastic.NewSearchSource().
		Suggester(termSuggester).
		FetchSource(false).
		TrackScores(true).
		Timeout("100ms")

	searchResult, _ := es.client.Search().
		Index(index).
		SearchSource(query).
		Do(context.Background())

	suggestions := searchResult.Suggest["movie-suggestion"]

	for _, i := range suggestions {
		for _, j := range i.Options {
			*suggestKeyword = j.Text
			break
		}
	}

	util.CombineSplitWords(suggestKeyword)

	ch <- true
}

func (es *DefaultElasticsearchService) BuildMatchQuery(text string, queue *util.Queue, fields ...string) {
	for _, value := range fields {
		queue.Enqueue(elastic.NewMatchQuery(value, text))
	}
}

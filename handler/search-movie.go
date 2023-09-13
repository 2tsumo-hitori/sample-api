package handler

import (
	"github.com/2tsumo-hitori/sample-api/elasticsearch"
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/2tsumo-hitori/sample-api/util"
)

const (
	movieNmText     = "movieNm_text"
	movieNmEngToKor = "movieNm_eng2kor"
	movieNmKorToEng = "movieNm_kor2eng"
)

func SearchByKeyword(searchKeyword string, resp *[]model.SearchResponse, es elasticsearch.SearchService) {
	suggestKeyword := searchKeyword
	q := util.Queue{}
	ch := make(chan bool)

	go es.BuildSuggestQuery(&suggestKeyword, ch)

	_, s := util.InspectSpell(searchKeyword)

	es.BuildMatchQuery(s, &q, movieNmText, movieNmEngToKor, movieNmKorToEng)
	es.SendRequestToElastic(&q, resp)

	if len(*resp) != 0 {
		return
	}

	select {
	case <-ch:
		es.BuildMatchQuery(suggestKeyword, &q, movieNmText)
		es.SendRequestToElastic(&q, resp)
		close(ch)
	}
}

func AutoCompleteByKeyword(searchKeyword string, resp *[]model.SearchResponse, es elasticsearch.SearchService) {
	q := util.Queue{}

	q.Enqueue(es.QueryBuildByKeyword(searchKeyword))

	es.SendRequestToElastic(&q, resp)
}

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

type DefaultService struct {
	Es elasticsearch.SearchService
}

func (ds *DefaultService) SearchByKeyword(searchKeyword string, resp *[]model.SearchResponse) {
	suggestKeyword := searchKeyword
	q := util.Queue{}
	ch := make(chan bool)

	go ds.Es.BuildSuggestQuery(&suggestKeyword, ch)

	_, s := util.InspectSpell(searchKeyword)

	ds.Es.BuildMatchQuery(s, &q, movieNmText, movieNmEngToKor, movieNmKorToEng)
	ds.Es.SendRequestToElastic(&q, resp)

	if len(*resp) != 0 {
		return
	}

	select {
	case <-ch:
		ds.Es.BuildMatchQuery(suggestKeyword, &q, movieNmText)
		ds.Es.SendRequestToElastic(&q, resp)
		close(ch)
	}
}

func (ds *DefaultService) AutoCompleteByKeyword(searchKeyword string, resp *[]model.SearchResponse) {
	q := util.Queue{}

	q.Enqueue(ds.Es.QueryBuildByKeyword(searchKeyword))

	ds.Es.SendRequestToElastic(&q, resp)
}

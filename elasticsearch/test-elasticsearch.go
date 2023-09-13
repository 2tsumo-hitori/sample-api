package elasticsearch

import (
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/2tsumo-hitori/sample-api/util"
)

func (es *TestService) SendRequestToElastic(queryQueue util.Queue, resp *[]model.SearchResponse) {
	*resp = append(*resp, model.SearchResponse{MovieNm: "테스트 성공"})
}

func (es *TestService) QueryBuildByKeyword(searchKeyword string) interface{} {
	return ""
}

func (es *TestService) BuildSuggestQuery(suggestKeyword *string, ch chan bool) {
	ch <- true
}

func (es *TestService) BuildMatchQuery(text string, queue *util.Queue, fields ...string) {
}

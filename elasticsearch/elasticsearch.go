package elasticsearch

import (
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/2tsumo-hitori/sample-api/util"
)

type SearchService interface {
	BuildSuggestQuery(suggestKeyword *string, ch chan bool)
	BuildMatchQuery(searchKeyword string, q *util.Queue, fields ...string)
	SendRequestToElastic(q *util.Queue, resp *[]model.SearchResponse)
	QueryBuildByKeyword(searchKeyword string) interface{}
}

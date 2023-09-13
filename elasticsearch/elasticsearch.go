package elasticsearch

import (
	"github.com/2tsumo-hitori/sample-api/config/esclient"
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/2tsumo-hitori/sample-api/util"
	"github.com/olivere/elastic/v7"
)

type SearchService interface {
	BuildSuggestQuery(suggestKeyword *string, ch chan bool)
	BuildMatchQuery(searchKeyword string, q *util.Queue, fields ...string)
	SendRequestToElastic(q *util.Queue, resp *[]model.SearchResponse)
	QueryBuildByKeyword(searchKeyword string) interface{}
}

type DefaultElasticsearchService struct {
	client *elastic.Client
}

func NewDefaultElasticsearchService() SearchService {
	return &DefaultElasticsearchService{
		client: esclient.Client(),
	}
}

type TestService struct{}

func NewTestService() SearchService {
	return &TestService{}
}

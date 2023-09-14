package handler_test

import (
	"fmt"
	"github.com/2tsumo-hitori/sample-api/elasticsearch"
	"github.com/2tsumo-hitori/sample-api/handler"
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var esService = handler.DefaultService{Es: elasticsearch.NewTestService()}

func TestBuildSuggestQuery(t *testing.T) {
	searchKeyword := "안녕"
	var resp []model.SearchResponse
	esService.SearchByKeyword(searchKeyword, &resp)

	fmt.Println(resp[0].MovieNm)

	assert.True(t, len(resp) != 0)
}

package handler_test

import (
	"fmt"
	"github.com/2tsumo-hitori/sample-api/elasticsearch"
	"github.com/2tsumo-hitori/sample-api/handler"
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildSuggestQuery(t *testing.T) {
	searchKeyword := "안녕"
	var resp []model.SearchResponse
	handler.SearchByKeyword(searchKeyword, &resp, elasticsearch.NewTestService())

	fmt.Println(resp[0].MovieNm)

	assert.True(t, len(resp) != 0)
}

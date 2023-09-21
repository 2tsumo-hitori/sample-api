package handler_test

import (
	"github.com/2tsumo-hitori/sample-api/elasticsearch"
	"github.com/2tsumo-hitori/sample-api/handler"
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var esService = handler.DefaultService{Es: elasticsearch.NewTestService()}

func TestSearchByKeyword(t *testing.T) {
	searchKeyword := "안녕"
	var resp []model.SearchResponse
	esService.SearchByKeyword(searchKeyword, &resp)

	assert.True(t, len(resp) != 0)
}

func TestAutoCompleteByKeyword(t *testing.T) {
	searchKeyword := "ㅎㅂㄹㄱ"
	var resp []model.SearchResponse

	esService.AutoCompleteByKeyword(searchKeyword, &resp)

	assert.True(t, strings.Contains(resp[0].MovieNm, "테스트 성공"))
}

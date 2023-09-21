package controller

import (
	"github.com/2tsumo-hitori/sample-api/elasticsearch/olivere"
	"github.com/2tsumo-hitori/sample-api/handler"
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/2tsumo-hitori/sample-api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 의존성 주입
var esService handler.DefaultService

func init() {
	esService = handler.DefaultService{Es: olivere.NewDefaultElasticsearchService()}
}

// MovieSearch 함수는 영화 검색을 제공합니다.
// @Summary 영화 검색
// @Description 검색 키워드에 해당되는 영화 목록을 제공합니다.
// @ID MovieSearch
// @Accept  json
// @Produce  json
// @Param request body model.MovieRequest true "영화 검색 요청 정보"
// @Success 200 {array} model.AutoCompleteResponse "검색된 영화 목록"
// @Router /es/search [post]
func (controller *Controller) MovieSearch(c *gin.Context) {
	var requestBody model.MovieRequest
	var movies []model.SearchResponse

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		panic(err)
	}

	esService.SearchByKeyword(requestBody.MovieNm, &movies)

	c.JSON(http.StatusOK, response.NewResponse(movies))
}

// AutoCompleteSearch 함수는 단어/초성 기반의 영화 자동 완성 검색을 제공합니다.
// @Summary 영화 자동 완성 검색
// @Description keydown을 감지할 때 마다 영화 키워드를 제공합니다.
// @ID autoCompleteSearch
// @Accept  json
// @Produce  json
// @Param request body model.MovieRequest true "영화 검색 요청 정보"
// @Success 200 {array} model.AutoCompleteResponse "검색된 영화 목록"
// @Router /es/ac [post]
func (controller *Controller) AutoCompleteSearch(c *gin.Context) {
	var requestBody model.MovieRequest
	var movies []model.SearchResponse

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		panic(err)
	}

	esService.AutoCompleteByKeyword(requestBody.MovieNm, &movies)

	c.JSON(http.StatusOK, response.NewResponse(movies))
}

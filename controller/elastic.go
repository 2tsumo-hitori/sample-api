package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-api/handler"
	"sample-api/model"
	"sample-api/response"
)

func (controller *Controller) AutoCompleteSearch(c *gin.Context) {
	var requestBody model.MovieRequest
	var movies []model.AutoCompleteResponse

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		panic(err)
	}

	handler.AutoComplete(requestBody.MovieNm, &movies)

	c.JSON(http.StatusOK, response.NewResponse().GetResponse(movies))
}

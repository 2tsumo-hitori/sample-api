package controller

import (
	"github.com/2tsumo-hitori/sample-api/model"
	"github.com/2tsumo-hitori/sample-api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) AlbumList(c *gin.Context) {
	c.JSON(http.StatusOK, response.NewResponse(model.Albums()))
}

func (controller *Controller) AddAlbum(c *gin.Context) {
	newAlbum := model.Album{}

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "잘못된 요청입니다."})
	} else {
		model.SetAlbum(newAlbum)

		c.JSON(http.StatusCreated, response.NewResponse(newAlbum))
	}
}

func (controller *Controller) Album(c *gin.Context) {
	id := c.Param("id")

	if _, exists := model.Albums()[id]; exists {
		c.JSON(http.StatusOK, response.NewResponse(model.Albums()[id]))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "잘못된 요청입니다."})
	}
}

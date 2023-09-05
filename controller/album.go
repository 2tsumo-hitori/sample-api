package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-api/model"
	"sample-api/response"
)

func (controller *Controller) AlbumList(c *gin.Context) {
	c.JSON(http.StatusOK, response.NewResponse().GetResponse(model.Albums()))
}

func (controller *Controller) AddAlbum(c *gin.Context) {
	newAlbum := model.Album{}

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "잘못된 요청입니다."})
	} else {
		model.SetAlbum(newAlbum)

		c.JSON(http.StatusCreated, response.NewResponse().GetResponse(newAlbum))
	}
}

func (controller *Controller) Album(c *gin.Context) {
	id := c.Param("id")

	if _, exists := model.Albums()[id]; exists {
		c.JSON(http.StatusOK, response.NewResponse().GetResponse(model.Albums()[id]))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "잘못된 요청입니다."})
	}
}

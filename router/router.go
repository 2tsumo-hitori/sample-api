package router

import (
	"github.com/gin-gonic/gin"
	"sample-api/controller"
)

func InitHandler() *gin.Engine {
	r := gin.Default()

	c := controller.NewController()

	r.GET("/albums", c.AlbumList)
	r.POST("/albums", c.AddAlbum)
	r.GET("/albums/:id", c.Album)

	return r
}

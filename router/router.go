package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"sample-api/controller"
	_ "sample-api/docs"
)

// @title           검색 API
// @version         1.0
// @description     검색 API OJT
// @host localhost:8080
// @BasePath /

func InitRouter() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	c := controller.NewController()

	r.GET("/albums", c.AlbumList)
	r.POST("/albums", c.AddAlbum)
	r.GET("/albums/:id", c.Album)

	r.POST("/es/search", c.MovieSearch)
	r.POST("/es/ac", c.AutoCompleteSearch)

	r.Run(":8080")
}

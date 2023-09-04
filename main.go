package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Data   interface{} `json:"album"`
	Status string      `json:"200"`
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = make(map[string]interface{})

func init() {
	albums["1"] = album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}
	albums["2"] = album{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99}
	albums["3"] = album{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99}
}

func setResponse(obj interface{}) response {
	var resp response

	resp.Data = obj
	resp.Status = "200"

	return resp
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/albums", func(c *gin.Context) {
		var resp response

		resp.Data = albums
		resp.Status = "200"

		c.JSON(http.StatusOK, resp)
	})

	r.POST("/albums", func(c *gin.Context) {
		var newAlbum album

		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}

		albums[newAlbum.ID] = newAlbum

		c.JSON(http.StatusCreated, newAlbum)
	})

	r.GET("/albums/:id", func(c *gin.Context) {
		if id := c.Param("id"); albums[id] != nil {
			c.JSON(http.StatusOK, setResponse(albums[id]))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "잘못된 요청입니다."})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

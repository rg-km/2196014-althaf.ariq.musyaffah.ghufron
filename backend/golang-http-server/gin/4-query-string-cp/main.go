package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// TODO: answer here
	router.GET("/movie", func(c *gin.Context) {
		genre := c.Query("genre")
		country := c.Query("country")
		if genre != "" && country != "" {
			//GET request method with genre and country query param
			c.String(http.StatusOK, "Here result of query of movie with genre %s in country %s", genre, country)
		} else if country != "" {
			//GET request method with country query param
			c.String(http.StatusOK, "Here result of query of movie with genre general in country %s", country)
		} else if genre != "" {
			//GET request method with genre query param
			c.String(http.StatusOK, "Here result of query of movie with genre %s", genre)
		} else {
			//GET request method without query param
			c.String(http.StatusOK, "Here result of query of movie with genre general")
		}
	})

	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": movies,
	}) // TODO: replace this
}

var MovieGetHandler = func(c *gin.Context) {
	id := c.Param("id")
	//convert id to int
	idInt, _ := strconv.Atoi(id)

	if idInt > 4 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movies[idInt],
	})
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	router.GET("/movie/list", MovieListHandler)
	router.GET("/movie/get/:id", MovieGetHandler)

	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}

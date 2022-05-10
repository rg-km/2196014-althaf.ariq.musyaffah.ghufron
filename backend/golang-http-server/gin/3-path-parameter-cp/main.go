package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Country string
	Age     int
}

var data = map[int]User{
	1: {
		Name:    "Roger",
		Country: "United States",
		Age:     70,
	},
	2: {
		Name:    "Tony",
		Country: "United States",
		Age:     40,
	},
	3: {
		Name:    "Asri",
		Country: "Indonesia",
		Age:     30,
	},
}

func ProfileHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		user, ok := data[idInt]
		if !ok {
			c.String(http.StatusNotFound, "data not found")
			return
		}

		c.String(http.StatusOK, "Name: %s, Country: %s, Age: %d", user.Name, user.Country, user.Age)

	}

}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/profile/:id", ProfileHandler())
	return r
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}

package main

import (
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
		idInt, _ := strconv.Atoi(id)
		if val, ok := data[int(idInt)]; !ok {
			c.String(http.StatusNotFound, "data not found")
		} else {
			c.String(http.StatusOK, "Name: %s, Country: %s, Age: %d", val.Name, val.Country, val.Age)
		}
	} // TODO: replace this
}

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/profile/:id", ProfileHandler())

	return router // TODO: replace this
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albuns = []album{
	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbuns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albuns)
}

func main (){
	router := gin.Default()
	router.GET("/albuns", getAlbuns)

	router.Run("localhost:3000")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albuns = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func postAlbuns (c *gin.Context){
	var newAlbun album

	if err := c.BindJSON(&newAlbun); err != nil{
		return
	}

	albuns = append(albuns, newAlbun)
	c.IndentedJSON(http.StatusCreated, newAlbun)
}

func getAlbuns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albuns)
}

func getAlbunsById (c *gin.Context){
	
	id := c.Param("id")

	for _, a := range albuns {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album nao funciona"})
}

func main (){
	router := gin.Default()
	router.GET("/albuns", getAlbuns)
	router.GET("/albuns/:id", getAlbunsById)
	router.POST("albuns", postAlbuns)


	router.Run("localhost:3000")
}

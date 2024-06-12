package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func returnDefinition(c *gin.Context) {
	searchWord := c.Param("searchWord")
	var dictionary map[int][]Word = loadDictionary()
	word := handleCollision(dictionary, searchWord)

	c.IndentedJSON(http.StatusOK, word.definition)
}
func main() {
	router := gin.Default()
	router.GET("/dictionary/:searchWord", returnDefinition)
	router.Run("localhost:8080")
}

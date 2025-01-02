package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Reciper struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}

func NewRecipeHandler(c *gin.Context) {
	
}

func main () {
	router := gin.Default()
	router.POST("/", NewRecipeHandler)
	router.Run()
	
}
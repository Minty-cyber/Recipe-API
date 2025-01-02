package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Recipe struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}
var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}
	


func NewRecipeHandler(c *gin.Context) {
	
}

func main () {
	router := gin.Default()
	router.POST("/", NewRecipeHandler)
	router.Run()
	
}
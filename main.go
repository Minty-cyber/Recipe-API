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
}

func main () {
	router := gin.Default()
	router.Run()
	
}
package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Reciper struct {
	ID string `json:"id"`
}

func main () {
	router := gin.Default()
	router.Run()
	
}
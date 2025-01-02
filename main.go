package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Reciper struct {
	
}

func main () {
	router := gin.Default()
	router.Run()
	
}
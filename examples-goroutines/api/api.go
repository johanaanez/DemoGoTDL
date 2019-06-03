package main

import (
	"github.com/gin-gonic/gin"
)

// main routine
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/bank")
	{
		v1.POST("/buy", makeBuy)
		v1.POST("/buyc", makeBuyConcurrent)
	}
	router.Run()
}

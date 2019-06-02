package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"math/rand"
	)

// ----------------------------------------

// SERVER MOCKS

// Simulates server 1

func getExchangeRate1(rate chan<- float64) {

	n := rand.Intn(3)

	time.Sleep(time.Duration(n) * time.Millisecond)
	rate <- 1.0 / 45.0
}

// Simulates server 2

func getExchangeRate2(rate chan<- float64 ) {

	n := rand.Intn(3)

	time.Sleep(time.Duration(n) * time.Millisecond)
	rate <- 1.0 / 45.0
}

// Simulates server 3

func getExchangeRate3(rate chan<- float64) {

	n := rand.Intn(3)

	time.Sleep(time.Duration(n) * time.Millisecond)
	rate <- 1.0 / 45.0
}

// ----------------------------------------

func getDollarExchangeRate(c* gin.Context) {

	exchangeRateChannel1 := make(chan float64)
	exchangeRateChannel2 := make(chan float64)
	exchangeRateChannel3 := make(chan float64)

	go getExchangeRate1(exchangeRateChannel1)
	go getExchangeRate2(exchangeRateChannel2)
	go getExchangeRate3(exchangeRateChannel3)

	var exchangeRate float64

	select {

	case response := <-exchangeRateChannel1:
		fmt.Println("Server 1 responded first")
		exchangeRate = response
	case response := <-exchangeRateChannel2:
		fmt.Println("Server 2 responded first")
		exchangeRate = response
	case response := <-exchangeRateChannel3:
		fmt.Println("Server 3 responded first")
		exchangeRate = response
	}

	c.JSON(200, gin.H{
		"dollar-exchange-rate": exchangeRate,
	})
}

func main() {
	
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Index",
		})
	})

	router.GET("/dollarExchangeRate", getDollarExchangeRate)

	router.Run(":8000")
}
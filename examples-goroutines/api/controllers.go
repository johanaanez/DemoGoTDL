package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResultDto struct {
	Valid   bool
	Message string
	Balance int
}

// POST body type
type MakeBuyDto struct {
	AccountNumber int `json:"AccountNumber" binding:"required"`
	Amount        int `json:"Amount" binding:"required"`
}

//handlers for enpoints
func makeBuy(c *gin.Context) {
	var data MakeBuyDto
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resultDto = buyForeingCurrency(data.Amount, data.AccountNumber)

	c.JSON(http.StatusOK, resultDto)
}

func makeBuyConcurrent(c *gin.Context) {
	var data MakeBuyDto
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resultDto = buyForeingCurrencyConcurrent(data.Amount, data.AccountNumber)

	c.JSON(http.StatusOK, resultDto)
}

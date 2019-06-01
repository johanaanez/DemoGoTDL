package main

import (
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const qtyServices = 3

type Result struct {
	ServiceId int
	Value     int
	Valid     bool
}

type ResultDto struct {
	Valid   bool
	Message string
}

// todoModel describes a todoModel type
type MakeBuyDto struct {
	AccountNumber int `json:"AccountNumber" binding:"required"`
	Amount        int `json:"Amount" binding:"required"`
}

func delay() {
	time.Sleep(1 * time.Second)
}

func getValue(bottom int, top int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(int(top-bottom)) + bottom
}

//------------------------
// three external services
func getExchangeRate() Result {
	delay()
	return Result{0, getValue(4000, 6000), true}
}

func validateOperationalTime() Result {
	delay()
	var hour = getValue(0, 2400)
	var valid = ((hour >= 900) && (hour <= 1300))
	return Result{1, 0, valid}
}

func getUserBalanceAccount() Result {
	delay()
	return Result{2, getValue(100000, 1000000), true}
}

//------------------------
func consumeServices() []Result {
	var results = []Result{}
	results = append(results, getExchangeRate())
	results = append(results, validateOperationalTime())
	results = append(results, getUserBalanceAccount())
	return results
}
func confirmOperation(accountNumber int, paidAmount int) bool {
	return true
}

func validateOperation(amount int, accountNumber int, serviceResults []Result) bool {
	var isValidOperationalTime = serviceResults[1].Valid
	var exchangeRate = serviceResults[0].Value
	var accountBalance = serviceResults[2].Value

	var paidAmount = amount * exchangeRate

	if (isValidOperationalTime) && (accountBalance >= paidAmount) {
		return confirmOperation(accountNumber, paidAmount)
	} else {
		return false
	}
}

func buyForeingCurrency(amount int, accountNumber int) ResultDto {
	var valid = validateOperation(amount, accountNumber, consumeServices())

	if valid {
		return ResultDto{valid, "Operation Confirmed"}
	} else {
		return ResultDto{valid, "Operation Invalid"}
	}
}

//--------------------------------
//functions for concurrent process
func callServiceAsync(service func() Result, results chan<- Result, wg *sync.WaitGroup) {
	results <- service()
	wg.Done()
}

func consumeServicesConcurrent() []Result {
	var results = make([]Result, qtyServices)
	resultsChan := make(chan Result, qtyServices)

	var w sync.WaitGroup
	w.Add(qtyServices)

	go callServiceAsync(getExchangeRate, resultsChan, &w)
	go callServiceAsync(validateOperationalTime, resultsChan, &w)
	go callServiceAsync(getUserBalanceAccount, resultsChan, &w)

	w.Wait()

	for i := 0; i <= qtyServices-1; i++ {
		r := <-resultsChan
		results[r.ServiceId] = r
	}

	close(resultsChan)
	return results
}

func buyForeingCurrencyConcurrent(amount int, accountNumber int) ResultDto {
	var valid = validateOperation(amount, accountNumber, consumeServicesConcurrent())
	if valid {
		return ResultDto{valid, "Operation Confirmed"}
	} else {
		return ResultDto{valid, "Operation Invalid"}
	}
}

//--------------------------------
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

//--------------------------------
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

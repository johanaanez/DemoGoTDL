package main

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
)

const qtyServices = 3

var userBalances = []int {2000, 1000, 2500, 5000}

var dolarValue = []int {46, 45, 48}

var mutex	sync.Mutex

type Result struct {
	ServiceId int
	Value     int
	Valid     bool
}

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

func delay() {
	time.Sleep(1 * time.Second)
}

func getValue(bottom int, top int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(int(top-bottom)) + bottom
}

//------------------------
// three external services
func getExchangeRate(timeId int) Result {
	delay()
	return Result{0, dolarValue[timeId], true}
}

func validateOperationalTime(hour int) Result {
	delay()
	var valid = ((hour >= 10) && (hour <= 2300))
	return Result{1, 0, valid}
}

func getUserBalanceAccount(userId int) Result {
	delay()
	return Result{2, userBalances[userId], true}
}

//------------------------
func consumeServices(userId int) []Result {
	var results = []Result{}
	results = append(results, getExchangeRate(getValue(0, 2)))
	results = append(results, validateOperationalTime(getValue(0, 2400)))
	results = append(results, getUserBalanceAccount(userId))
	return results
}

func confirmOperation(accountNumber int, paidAmount int) (bool, int) {
	mutex.Lock()
	userBalances[accountNumber] -= paidAmount
	mutex.Unlock()
	return true, userBalances[accountNumber]
}

func validateOperation(amount int, accountNumber int, serviceResults []Result) (bool, int) {
	var isValidOperationalTime = serviceResults[1].Valid
	var exchangeRate = serviceResults[0].Value
	var accountBalance = serviceResults[2].Value

	var paidAmount = amount * exchangeRate

	if (isValidOperationalTime) && (accountBalance >= paidAmount) {
		return confirmOperation(accountNumber, paidAmount)
	} else {
		return false, accountBalance
	}
}

func buyForeingCurrency(amount int, accountNumber int) ResultDto {
	var valid, balance = validateOperation(amount, accountNumber, consumeServices(accountNumber))

	if valid {
		return ResultDto{valid, "Operation Confirmed", balance}
	} else {
		return ResultDto{valid, "Operation Invalid", balance}
	}
}

//--------------------------------
//functions for concurrent process
func callServiceAsync(service func(a int) Result, results chan<- Result, wg *sync.WaitGroup, value int) {
	results <- service(value)
	wg.Done()
}

func consumeServicesConcurrent(userId int) []Result {
	var results = make([]Result, qtyServices)
	resultsChan := make(chan Result, qtyServices)

	var w sync.WaitGroup
	w.Add(qtyServices)

	go callServiceAsync(getExchangeRate, resultsChan, &w, getValue(0, 2))
	go callServiceAsync(validateOperationalTime, resultsChan, &w, getValue(0, 2400))
	go callServiceAsync(getUserBalanceAccount, resultsChan, &w, userId)

	w.Wait()

	for i := 0; i <= qtyServices-1; i++ {
		r := <-resultsChan
		results[r.ServiceId] = r
	}

	close(resultsChan)
	return results
}

func buyForeingCurrencyConcurrent(amount int, accountNumber int) ResultDto {
	var valid, balance = validateOperation(amount, accountNumber, consumeServicesConcurrent(accountNumber))
	if valid {
		return ResultDto{valid, "Operation Confirmed", balance}
	} else {
		return ResultDto{valid, "Operation Invalid", balance}
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

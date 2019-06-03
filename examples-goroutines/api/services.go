package main

import (
	"math/rand"
	"time"
)

var userBalances = []int{2000, 1000, 2500, 5000}

var dolarValue = []int{46, 45, 48, 49}

type Result struct {
	ServiceId int
	Value     int
	Valid     bool
}

func delay() {
	time.Sleep(10 * time.Millisecond)
}

func getValue(bottom int, top int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(int(top-bottom)) + bottom
}

// simple mock for get exchangeRate
func getExchangeRate0(timeId int) Result {
	delay()
	return Result{0, dolarValue[timeId], true}
}

// multiple services for concurrent use
// in order to get exchangeRate
func getExchangeRate1(timeId int, rate chan<- int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	rate <- dolarValue[timeId]
}

func getExchangeRate2(timeId int, rate chan<- int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	rate <- dolarValue[timeId]
}

func getExchangeRate3(timeId int, rate chan<- int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	rate <- dolarValue[timeId]
}

// three external services mocks
func getExchangeRate(timeId int) Result {
	delay()
	exchangeRateChannel1 := make(chan int)
	exchangeRateChannel2 := make(chan int)
	exchangeRateChannel3 := make(chan int)

	go getExchangeRate1(timeId, exchangeRateChannel1)
	go getExchangeRate2(timeId, exchangeRateChannel2)
	go getExchangeRate3(timeId, exchangeRateChannel3)

	var exchangeRate int

	select {

	case response := <-exchangeRateChannel1:
		exchangeRate = response
	case response := <-exchangeRateChannel2:
		exchangeRate = response
	case response := <-exchangeRateChannel3:
		exchangeRate = response
	}
	return Result{0, exchangeRate, true}
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

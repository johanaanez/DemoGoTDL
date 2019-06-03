package main

import (
	"sync"
)

const qtyServices = 3

//functions for concurrent process
func callServiceAsync(service func(a int) Result, results chan<- Result, wg *sync.WaitGroup, value int) {
	results <- service(value)
	wg.Done()
}

func consumeServicesConcurrent(userId int) []Result {
	var results = make([]Result, qtyServices)
	resultsChan := make(chan Result, qtyServices)
	defer close(resultsChan)

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

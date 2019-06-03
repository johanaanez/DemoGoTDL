package main

import "sync"

var mutex sync.Mutex

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

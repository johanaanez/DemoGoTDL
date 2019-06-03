package main

import "errors"

func consumeServices(userId int) []Result {
	var results = []Result{}
	results = append(results, getExchangeRate0(getValue(0, 2)))
	results = append(results, validateOperationalTime(getValue(0, 24)))
	results = append(results, getUserBalanceAccount(userId))
	return results
}

func buyForeingCurrency(amount int, accountNumber int) (result ResultDto, err error) {

	defer func() {
		if panicError := recover(); panicError != nil {
			err = errors.New("crital error ocurred")
		}
	}()

	var valid, balance = validateOperation(amount, accountNumber, consumeServices(accountNumber))

	if valid {
		return ResultDto{valid, "Operation Confirmed", balance}, nil
	} else {
		return ResultDto{valid, "Operation Invalid", balance}, nil
	}

}

package main

import (
	"fmt"
	"time"
)


func isHomeBankingAvailable(now chan time.Time) {
	date := <-now
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday || date.Hour() < 10 || date.Hour() > 14 {
		fmt.Println("Bank is close")
	} else {
		fmt.Println("Bank is open")
	}
}

func getDollarsToBuy(value chan int) {
	dollar := 46
	convertedValue := <-value / dollar
	fmt.Printf("Dollars : %d \n", convertedValue)
}


func main() {
	fmt.Println("Start")
	client := Client{"23939477654", 46000}
	channel := make(chan int)
	clientCh := make(chan Client)
	now := make(chan time.Time)

	go isHomeBankingAvailable(now)
	now <- time.Now()
	time.Sleep(1 * time.Millisecond)

	go getDollarsToBuy(channel)
	channel <- 46000

	go hasAvailableBalance(clientCh)
	clientCh <- client
	time.Sleep(1 * time.Millisecond)

	fmt.Println("End")
}

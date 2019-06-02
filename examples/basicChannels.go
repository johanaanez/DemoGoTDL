package main

import (
	"fmt"
	"time"
)

func isHomeBankingAvailable() {
	now := time.Now()
	if (now.Weekday() == time.Saturday || now.Weekday() == time.Sunday || now.Hour() < 10 || now.Hour()>14 ) {
		fmt.Println("Closed")
	} else {
		fmt.Println("Open")
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

	go isHomeBankingAvailable()
	time.Sleep(1 * time.Millisecond)

	go getDollarsToBuy(channel)
	channel <- 46000

	go hasAvailableBalance(clientCh)
	clientCh <- client

	fmt.Println("End")
}

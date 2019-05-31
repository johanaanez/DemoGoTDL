package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func worker(qty int, results chan<- int) {
    results <- getRandomNumber(qty)
}

func getRandomNumber(qty int) int {
	time.Sleep(3 * time.Second)
	rand.Seed(time.Now().UnixNano())
    max := math.Pow10(qty)
	return rand.Intn(int(max))
}

func mainSecuential(){
	list := []int{4, 5, 7}
	result := []int {}
	for _, qty := range list {
		result = append(result, getRandomNumber(qty))
	}

	fmt.Println(result)
}

func mainConcurrent(){
	list := []int{4, 5, 7}
	numberOfCalculations := len(list)
	resultList := []int {}
	resultsChan := make(chan int, numberOfCalculations)

	for _, qty := range list {
		go worker(qty, resultsChan)
	}

	for i := 0; i <= numberOfCalculations-1; i++ {
		r := <-resultsChan
		resultList = append(resultList, r)
	}

	close(resultsChan)
	
	fmt.Println(resultList)

}

func main(){
	//mainSecuential()
	mainConcurrent()
}
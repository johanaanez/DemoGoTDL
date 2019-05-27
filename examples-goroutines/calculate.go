package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func getRandomNumber(qty int) int {
	time.Sleep(3 * time.Second)
	rand.Seed(time.Now().UnixNano())
    max := math.Pow10(qty)
	return rand.Intn(int(max))
}

func main(){
	list := []int{4, 5, 7}
	result := []int {}
	for _, qty := range list {
		result = append(result, getRandomNumber(qty))
	}

	fmt.Println(result)

}
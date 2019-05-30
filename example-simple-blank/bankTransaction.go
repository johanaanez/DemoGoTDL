package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex	sync.Mutex
)

func init() {
	balance = 1000
}

func deposit(index int, value int, wg *sync.WaitGroup) {
	//mutex.Lock()
	balance += value
	fmt.Printf("%d) Deposito: %d. Balance en la cuenta: %d\n", index, value, balance)
	wg.Done()
	//mutex.Unlock()
}

func extract(index int, value int, wg *sync.WaitGroup) {
	//mutex.Lock()
	balance -= value
	fmt.Printf("%d) Extraccion: %d. Balance en la cuenta: %d\n", index, value, balance)
	wg.Done()
	//mutex.Unlock()
}

func main() {
	fmt.Println("¡Bienvenidos a GObank!")

	var w sync.WaitGroup
	for i := 1; i <= 100; i++ {
		w.Add(2)        
		go extract(i, 200, &w)
		go deposit(i, 200, &w)
	}		
	w.Wait()

	fmt.Printf("Balance final: %d\n", balance)
}
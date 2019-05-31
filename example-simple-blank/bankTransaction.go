package main

import (
	"fmt"
	"sync"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"log"
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

func depositHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	value, _ := strconv.Atoi(params["value"])

	balance += value
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Extraccion: %d. Balance en la cuenta: %d\n", value, balance)
}

func extractHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	value, _ := strconv.Atoi(params["value"])

	balance -= value
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Extraccion: %d. Balance en la cuenta: %d\n", value, balance)
}

func concurrentActionHandler(w http.ResponseWriter, r *http.Request) {
	var waitG sync.WaitGroup
	params := mux.Vars(r)
	cant, _ := strconv.Atoi(params["cantActions"])

	for i := 1; i <= cant; i++ {
		waitG.Add(2)
    go extract(i, 200, &waitG)
		go deposit(i, 200, &waitG)
	}

	waitG.Wait()
}

func main() {
	fmt.Println("¡Bienvenidos a GObank!")

	r := mux.NewRouter()
	r.HandleFunc("/deposit/{value}", depositHandler).Methods("GET")
	r.HandleFunc("/deposit/{value}", extractHandler).Methods("GET")
	// r.HandleFunc("/actions/{cantActions}", actionHandler).Methods("GET")
	r.HandleFunc("/concurrentActions/{cantActions}", concurrentActionHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("locahost:8000", r))
}
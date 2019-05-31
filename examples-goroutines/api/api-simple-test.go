package main

import (
	"time"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"math/rand"
)

const qtyServices = 3

type Result struct {
	ServiceId int
	Value  	int
}


func delay() {
	time.Sleep(1 * time.Second)
}

func consumeService() int {
	delay()
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(int(4000))
}

func consumeServices() []Result {
	var results = []Result{}
	for i := 0; i < 3; i++ {
		results = append(results, Result{i, consumeService()})
	}
	return results
}




func getBestResult(results []Result) Result{
	var bestResult Result

	for i, e := range results {
		if i==0 || e.Value < bestResult.Value {
			bestResult = e
		}
	}
	return bestResult
}

func getResult() Result{
	return getBestResult(consumeServices())
}
//--------------------------------
//functions for concurrent process
func worker(serviceId int, results chan<- Result) {
    results <- Result{serviceId, consumeService()}
}

func consumeServicesConcurrent() []Result {
	var results = []Result{}
	resultsChan := make(chan Result, qtyServices)

	for i := 0; i <= qtyServices-1; i++ {
		go worker(i, resultsChan)
	}

	for i := 0; i <= qtyServices-1; i++ {
		r := <-resultsChan
		results = append(results, r)
	}

	close(resultsChan)
	return results
}

func getResultConcurrent() Result{
	return getBestResult(consumeServicesConcurrent())
}
//--------------------------------
//handlers
func getResource(w http.ResponseWriter, r *http.Request) {
	
	var result = getResult()

	json.NewEncoder(w).Encode(result)
}

func getResourceConcurrent(w http.ResponseWriter, r *http.Request) {
	
	var result = getResultConcurrent()

	json.NewEncoder(w).Encode(result)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getResource).Methods("GET")
	router.HandleFunc("/concurrent", getResourceConcurrent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
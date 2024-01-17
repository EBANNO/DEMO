package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	//start http server
	http.HandleFunc("/migrate", migrate_v1)
	http.ListenAndServe(":3000", nil)
}

// get product list and process each product
func migrate_v1(w http.ResponseWriter, r *http.Request) {
	//get products
	products := getProducts()
	for _, product := range products {
		//process product
		processProduct(product)
	}

	//return response
	response, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
func migrate_v2(w http.ResponseWriter, r *http.Request) {
	//create wait group
	wg := sync.WaitGroup{}

	//get products
	products := getProducts()
	wg.Add(len(products)) //set counter
	for _, product := range products {
		//process product
		go func(productPorcess *Product) {
			processProduct(productPorcess)
			wg.Done() //decrement counter
		}(product)
	}
	wg.Wait() //wait for all goroutines to finish

	//return response
	response, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// get products
func getProducts() []*Product {
	//get random number of products
	products := []*Product{}
	for id := 0; id < 1000; id++ {
		products = append(products, &Product{
			ID: strconv.Itoa(id),
		})
	}
	time.Sleep(time.Duration(1 * time.Second)) //simulate delay
	return products
}

// process transaction and change status
func processProduct(product *Product) {
	time.Sleep(time.Duration(1 * time.Second)) //simulate delay
	fmt.Printf("Process product: %s\n", product.ID)
	//change status to success or fail randomly
	if rand.Intn(2) == 1 {
		product.Status = "success"
	} else {
		product.Status = "fail"
	}
}

type Product struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

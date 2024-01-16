package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	//start http server
	http.HandleFunc("/notify", notify)
	http.ListenAndServe(":3000", nil)
}

func notify(w http.ResponseWriter, r *http.Request) {
	productChannel := make(chan Product)

	//get product
	go getProducts(productChannel)

	//notifications
	go emailNotify(productChannel)
	go smsNotify(productChannel)
	go eventNotify(productChannel)
}

// get products and put it on channel
func getProducts(productChannel chan<- Product) {
	//get random number of product
	for id := 0; id < 1000; id++ {
		product := Product{
			ID: strconv.Itoa(id),
		}
		fmt.Printf("Get product: %s\n", product.ID)
		time.Sleep(1 * time.Second) //simulate delay

		productChannel <- product
	}
	//close channel
	close(productChannel)
}

// email notification by channel
func emailNotify(productChannel <-chan Product) {
	for product := range productChannel {
		// TODO: implement notification
		fmt.Printf("Email notification product: %s\n", product.ID)
		time.Sleep(1 * time.Second) //simulate delay
	}
}

// sms notification by channel
func smsNotify(productChannel <-chan Product) {
	for product := range productChannel {
		// TODO: implement notification
		fmt.Printf("SMS notification product: %s\n", product.ID)
		time.Sleep(1 * time.Second) //simulate delay
	}
}

// event notification by channel
func eventNotify(productChannel <-chan Product) {
	for product := range productChannel {
		// TODO: implement notification
		fmt.Printf("Event notification product: %s\n", product.ID)
		time.Sleep(1 * time.Second) //simulate delay
	}
}

type Product struct {
	ID string
}

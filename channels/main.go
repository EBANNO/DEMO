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

// get product list and notify each product on email
func notify(w http.ResponseWriter, r *http.Request) {
	emailProductChannel := make(chan Product)
	smsProductChannel := make(chan Product)
	eventProductChannel := make(chan Product)

	//get product
	go getProducts(emailProductChannel, smsProductChannel, eventProductChannel)

	//notifications
	go emailNotify(emailProductChannel)
	go smsNotify(smsProductChannel)
	go eventNotify(eventProductChannel)
}

/*
getProducts: info -> channel
emailNotify: channel -> info
smsNotify: channel -> info
eventNotify: channel -> info
*/

// get products and put it on channel
func getProducts(emailProductChannel, smsProductChannel, eventProductChannel chan<- Product) {
	//get random number of product
	for id := 0; id < 1000; id++ {
		product := Product{
			ID: strconv.Itoa(id),
		}
		fmt.Printf("Get product: %s\n", product.ID)
		time.Sleep(1 * time.Second) //simulate delay

		emailProductChannel <- product
		smsProductChannel <- product
		eventProductChannel <- product
	}
	//close channel
	close(emailProductChannel)
	close(smsProductChannel)
	close(eventProductChannel)
}

// email notification by channel
func emailNotify(emailProductChannel <-chan Product) {
	for product := range emailProductChannel {
		// TODO: implement notification
		fmt.Printf("Email notification product: %s\n", product.ID)
		time.Sleep(1 * time.Second) //simulate delay
	}
}

// sms notification by channel
func smsNotify(eventProductChannel <-chan Product) {
	for product := range eventProductChannel {
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

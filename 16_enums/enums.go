package main

import "fmt"

type OrderStatus string

const (
	Received OrderStatus = "Received"
	Confirmed = "Confirmed"
	Prepared = "Prepared"
	Delivered = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Chaning order status to", status)
}

func main() {
	changeOrderStatus(Received)
}
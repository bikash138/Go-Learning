package main

import (
	"fmt"
	"time"
)

type customer struct {
	name string
	phone string
}

type order struct {
	id      string
	amount    float32
	status    string
	createdAt time.Time
	customer //Struct embedding
}

//While the values of the struct we need to pass the pointer
func (o *order) changeStatus(status string) {
	o.status = status
}

//For getting the values from the struct we need not to pass the pointer
func (o order) getAmount() float32 {
	return o.amount
}

//Creating a function to create objecs of order struct like ctor in oop
//We are returning the pointer to the created object
func newOrder(id string, amount float32, status string) *order {
	myOrder := order {
		id: id,
		amount: amount,
		status: status,
		customer: customer{
			name: "Bikash",
			phone: "123456789",
		},
	}
	return &myOrder
}

func main() {

	//So now we will call the newOrder function to create a new order
	myOrder1 := newOrder("2", 89.00, "Placed")

	fmt.Println(myOrder1)

	myOrder := order {
		id: "1",
		amount: 20.00,
		status: "Shipped",
	}
	myOrder.changeStatus("Confirmed")
	fmt.Println(myOrder.getAmount())

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder)

	//We can inline struct like:
	language := struct {
		name string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
package main

import "fmt"

//In go we need not to implement the the interfaces 
//Its automatically gets infered to the struct whose pay method looks same
type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment made using razorpay", amount)
}

type stripe struct{}

//To implement the paymenter interface write the pay method as it is in th interface
func (s stripe) pay(amount float32) {
	fmt.Println("Payment made using stripe", amount)
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

func main() {
	//We can pass any payment gateway which are implementing the paymenter interface
	// stripePaymentGw := stripe{}
	stripePaymentGw := razorpay{}
	newPayment := payment{
		gateway: stripePaymentGw,
	}
	newPayment.makePayment(100)
}
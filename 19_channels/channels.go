package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	fmt.Println("Processing number", <- numChan)
}

func sum(result chan int, num1 int, num2 int) {
	ans := num1 + num2
	result <- ans
}
func main() {
	result := make(chan int)
	go sum(result, 4, 8)
	res := <- result //blocking->Here we dont need to stop the function for the go routines to run
	fmt.Println(res)

	//Channels are used to send the data from one goroutine to another goroutine
	numChan := make(chan int)

	go processNum(numChan)

	numChan <- 5

	time.Sleep(time.Second * 2)
}
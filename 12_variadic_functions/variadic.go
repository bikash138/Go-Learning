package main

import "fmt"

//To use variadic funtion we use "..."
//Here the nums become slice
func sum (nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	return sum
}

func main() {
	//Variadic Function are the functions that can accept variable number of arguments
	//For exmple: fmt.Println where we can as many args as we want
	// fmt.Println(10,20,30, "SOmethig")
	answer := sum(12,10,20)
	fmt.Println(answer)
}
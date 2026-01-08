package main

import "fmt"

//We are receing the num variable as value so copy of it will be created
func changeNum(num int) {
	num = 5
	fmt.Println("Changed NUmber", num)
}

//So now we will take the input as reference means as pointer
func changeNumRef(num *int) {
	//Destructure the pointer to modify the value
	*num = 5
	fmt.Println("Changed NUmber", *num)
}

func main() {
	num := 1

	// changeNum(num)
	//Use & for passing the memory address
	changeNumRef(&num)
	//Here the num will not change because we had passed the num as value not by reference
	fmt.Println("In main num: ", num)
}
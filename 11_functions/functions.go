package main

import "fmt"

func add(a, b int) int {
	return a + b
}

//Returning multiple values -> Order of return values should be same
func getSomething() (string, int, string) {
	return "Bikash", 29, "golang"
}	

func main() {
	ans := add(4, 8)
	fmt.Println(ans)
	fmt.Println(getSomething())
}
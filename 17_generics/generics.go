package main

import "fmt"

//We can use genrics to pass our own data types

func printSlice[T string | int](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	nums := []int {1, 2, 3}
	// lang := []string {"cpp", "golang", "typescript"}
	printSlice(nums)
}
package main

import "fmt"

//Range is mainly used for doing iteration over data structures
func main () {
	nums := []int {4,3,6}

	sum := 0
	//i says the index of the slice
	for i, num:= range nums {
		sum = sum + num
		fmt.Println(sum, i)
	}

	//Using range in maps
	m := map[string]string {"name": "Bikash", "lname": "Shaw"}
	//This is the convention to give name k -> key and v -> value
	for k, v:= range m {
		fmt.Println(k , v)
	}

	//We can also use range in string
	c := "Bikash"

	//Here v will represent the unicode of letters
	//Range iterates over the runes
	for _, v:=range c {
		fmt.Println(v)
	}

}	
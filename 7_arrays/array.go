package main

import "fmt"

func main() {
	//in arrays the zeroed values get fille with the the data according to datatype
	//int -> 0, bool -> false, string -> ""
	var st [5]string
	st[2] = "golang"
	fmt.Println(len(st))
	fmt.Println(st)

	//2d Arrays
	nums := [2][2] int {{1,2}, {1,2}}
	fmt.Println(nums)

	//Normally we dont use array in go becuase for the arrays the sie should be predictable
	//We use Slices where the memory is allocated dynamically
}
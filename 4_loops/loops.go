package main

import "fmt"

//In go there is only for loop (while loop is no present here)

func main() {
	//Implementing while using for loop
	i := 1
	for i<=2 {
		fmt.Println(i)
		i++
	}

	//Infinite loop
	// for {
	// 	println("1")
	// }

	//Normal for loop
	for i:=0; i<7; i++ {
		
		//We can use break/continue in the loops
		if(i == 2) {
			//Will skip the printing of  i when it is 2
			continue
		}

		if(i == 4){
			//Loop will stop when i becomes 4
			break
		}
		fmt.Println(i)
	}

	//We can use modern way to represent loops
	//Range starts the loop from 0 -> 4 (where 4 is not included)
	for i:= range 4 {
		fmt.Println(i)
	}

}
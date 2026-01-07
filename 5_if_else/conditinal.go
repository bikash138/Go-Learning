package main

import "fmt"

func main() {
	age := 14

	if age < 14 {
		fmt.Println("Child")
	} else if age < 18 && age >= 14 {
		fmt.Println("Teenage")
	} else {
		fmt.Println("Adulu")
	}

	//We can create a variable inside the if statement which can be accessed inside its scope as well as else scope
	if isAdmin:=true; isAdmin {
		fmt.Printf("Yes access allowed!!")
	}else {
		fmt.Printf("is now allowed")
	}

	//GO does not have ternary operator
}
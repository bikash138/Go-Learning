package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
		case 1:
			fmt.Printf("One")
		case 2:
			fmt.Printf("two")
		case 3:
			fmt.Printf("three")
		default:
			fmt.Printf("other")
	}

	day := time.Now().Weekday()
	switch day{
	case time.Saturday, time.Sunday: 
		fmt.Printf("%s: It's a weekend\n", day)
	default:
		fmt.Printf("%s: Its a normla work day\n", day)
	}

	//Type switch
	whoAmI := func(i any) {
		switch v := i.(type) {
		case string:
			fmt.Println("String", v)
		case int:
			fmt.Println("integer", v)
		case float64:
			fmt.Println("float", v)
		default:
			fmt.Println("Other", v)
		}

	}

	whoAmI(8.9)
}
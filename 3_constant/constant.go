package main
import "fmt"

func main() {
	//This value cannote be changed 
	//The value of constant should be assigned here unlike var whose value can be assigned anywhere
	const name string = "golang"

	//Cannot assing other values to it
	// name = "something"

	//Grouping if multiple constants
	const(
		port = 7000
		host = "localhost"
	)

	fmt.Println(name, port)
}
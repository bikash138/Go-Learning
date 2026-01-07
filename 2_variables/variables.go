package main
import "fmt"

func main() {
	
	//Here we have to give the type while creating the variable
	var name string = "golang"

	//Apart from this it infers the type from the data
	var name2 = "python"

	//Shorthand Syntax
	name3 := "typescript"

	fmt.Println(name, name2, name3)
}
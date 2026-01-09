package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name: ", fileInfo.Name())

	//Reading file using buffer array
	buf := make([]byte, 12)
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	for i:=0; i<len(buf); i++{
		fmt.Println("data", d, string(buf[i]))
	}

	//Simper way to rea file
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//Creating File
	c, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c.WriteString("Hello guys")
	c.WriteString("Hello bruh")


}
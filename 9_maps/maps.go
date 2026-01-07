package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)

	m["language"] = "golang"
	m["system"] = "macos"

	fmt.Println(m["language"])
	//IMP: If we try to access the key which is not present in the map then it will return zero value
	fmt.Println(m["phone"])
	//we can use the len method to get the length of the map
	fmt.Println(len(m))

	a := map[string]int {"age": 14, "phone": 1234567}
	//In go we have convetion to use ok which gives boolen value
	//another thing is v which is returned as the value of the key
	v, ok := a["age"]
	fmt.Println(v)

	if ok {
		fmt.Println("All ok")
	} else{
		fmt.Println("Not Ok")
	}
	b := map[string]int {"age": 12, "phones": 234}
	//Just like slices package we have map package for maps data structure
	fmt.Print(maps.Equal(a,b))
}
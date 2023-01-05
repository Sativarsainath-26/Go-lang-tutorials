package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	name     string
	price    int
	platform string
	password string
	courses  []string
}

func main() {
	fmt.Println("welcome to a class on creating json data in golang")

	// In this video we will be seeing encoding of json
	// that means that we pass data and convert that data into valid json
	lcoCourses := []Course{
		{"ReactjsBootcamp", 2000, "youtube", "sai.com", []string{"webdevelopment", "developer", "tester"}},
	}

	//package this above data into json
	fmt.Println(lcoCourses)

	finalJson, err := json.MarshalIndent(lcoCourses, " ", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

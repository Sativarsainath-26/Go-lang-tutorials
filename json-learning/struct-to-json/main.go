package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Welcome to a clas on converting struct to json")

	sainath := User{"sainath", "sainath.email.com", 16, true}
	fmt.Println(sainath)
	fmt.Printf("sainath details are : %+v\n", sainath)

	marshal_struct, err := json.Marshal(sainath)
	if err == nil {
		fmt.Println("the value of marshal_struct is given by", marshal_struct)
		fmt.Println("the value of marshal_struct in string is given by", string(marshal_struct))
	}
}

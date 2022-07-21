package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to a class on struct in golang")

	//no inheritance in golang
	//there is no super or parent class concept.

	sainath := User{"sainath", "sainath.email.com", 16, true}
	fmt.Println(sainath)
	fmt.Printf("sainath details are : %+v\n", sainath)
	fmt.Printf("Name is  : %v", sainath.Name)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

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
	fmt.Println("welcome to a class on conversion of json to struct")
	user_info := `{"Name":"sainath","Email":"sainath.email.com","Age":16,"Status":true}`

	user_info_bytes := []byte(user_info)

	var student User
	err := json.Unmarshal(user_info_bytes, &student)
	if err != nil {
		fmt.Println("the value of error is", err)
	}
	fmt.Println(student.Age)
	fmt.Println(student.Name)
	fmt.Println(student)

	// unmarshalling a JSON array
	// to array type in Golang

	user_info1 := `[{"Name":"sainath","Email":"sainath.email.com","Age":16,"Status":true},
	              {"Name":"lassy","Email":"sai.email.com","Age":14,"status":false}]`

	user_info1_bytes := []byte(user_info1)

	// defining an array instance
	// of struct type
	var student1 []User

	err1 := json.Unmarshal(user_info1_bytes, &student1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(student1)
	for i := range student1 {
		fmt.Println(student1[i])
	}

}

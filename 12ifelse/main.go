package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on if-else in Golang")

	logincount := 10
	var result string

	if logincount < 10 {
		result = "Regulated user"
	} else if logincount > 10 {
		result = "unregulated user"
	} else {
		result = "exactly 10 logincount"

	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}

}

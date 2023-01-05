package main

import "fmt"

func main() {

	fmt.Println("welcome to a class on functions in Golang")
	greeter() // calling a function
	greeterTwo()

	/*	func greeterTwo() {
		fmt.Println("another function ")    we should not write a function in function
	}  */
	result := adder(3, 5)

	fmt.Println("result is :", result)

	proresult, mymessage := proAdder(2, 5, 8, 9)
	fmt.Println("the result of addition is :", proresult)
	fmt.Println("the result of addition is :", mymessage)
}

func proAdder(value ...int) (int, string) {
	total := 0 // when we don't know how many values are going to get added.

	for _, val := range value {
		total = total + val

	}
	return total, "proResult function "
}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}

func greeterTwo() {
	fmt.Println("another function")
}
func greeter() {
	fmt.Println("namastey from golang")
}

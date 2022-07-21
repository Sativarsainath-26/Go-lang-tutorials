package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointer")
	/* why pointers are came ?
	Ans : whenever you are passing variables to function or classes or objects so the copy of the variable is passed instead of passing
	that variable. so the pointers make sure that it stores the variable address and we pass that pointer it  make sures that
	 we are passing exact value  passing exact value not the copied value. */

	// var ptr *int //initializing a pointer

	// fmt.Println("value of pointer is : ", ptr)  >>>>>> output is value of pointer is : <nil>

	mynumber := 23
	var ptr = &mynumber // referncing means &

	fmt.Println("The value of the actual pointer is ", ptr)  //it gives the address of the mynumber
	fmt.Println("The value of the actual pointer is ", *ptr) //it gives the value of the mynumber

	*ptr = *ptr + 2
	fmt.Println("the value of mynumber is :", mynumber)

}

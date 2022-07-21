package main

import "fmt"

//user : = 500  >>>>>//cannot use short hand declaration operator outside the method or function this is not allowed
// var user:=500 >>>this is allowed
// var user int = 50 >>> this is allowed

const Logintoken string = "sainath" //In golang we keep as first letter as a capital for a public variable.

func main() {
	//create a variable
	var username string = "Sainath"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var islogin bool = false
	fmt.Println(islogin)
	fmt.Printf("variable is of type : %T \n", islogin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type : %T \n", smallval)

	var smallfloat float32 = 255.433
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type : %T \n", smallfloat)

	//we are going to create aliases

	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type : %T\n", anothervariable)

	//implicit type go lang compiler implicitly understood that this type is string.

	var website = "sainath.com"
	fmt.Println(website)
	fmt.Printf("variable is of type : %T\n", website)

	//short hand declaration without using var

	numberofUser := 4000
	fmt.Println(numberofUser)
	fmt.Printf("variable is of type : %T \n", numberofUser)

	fmt.Println(Logintoken)
	fmt.Printf("variable is of type : %T\n", Logintoken)
}

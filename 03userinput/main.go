package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := " welcome to user input"
	fmt.Println(welcome)
	/*using bufio package and os.stdin we are reading the input from the user */
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating of our pizza")

	//comma ok syntax or error ok syntax

	input, _ := reader.ReadString('\n')
	/*here it reads the string until new line. why we are using
	input, _ because the reader reads the
	string from standard input if it goes something wrong then it stores in underscore.
	if everything goes right then it stores in input */
	fmt.Println("thanks for reading", input)
	fmt.Printf("type of the input is : %T", input)

}

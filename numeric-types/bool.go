//A boolean value is a binary variable having true or false. It plays an important role in programming.

package main

import (
	"fmt"
)

var x bool

func main() {

	fmt.Println(x) // output: false. By default boolean variable is declared it takes value as false.
	x = true
	fmt.Println(x)
}

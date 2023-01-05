package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := "sainath"

	marshal_str, err := json.Marshal(str)
	if err != nil {
		panic(err)
	}
	// the marshal function stores in the form of slice of bytes
	fmt.Println("the value of string after converting to marshal is", marshal_str)
	//output is the value of string after converting to marshal is [34 115 97 105 110 97 116 104 34]
	//Here it stores the values in the form of byte means the value of each letter is in the ascii value.
	fmt.Println("the value of string in converstion bytes to string is", string(marshal_str))
	//output is the value of string in converstion bytes to string is "sainath"

	/*for j := range string(marshal_str) {
		fmt.Println("the value of j is given by", j)
		fmt.Printf("the type of j is %T ", j)

		Need to understand further on this
	}*/
}

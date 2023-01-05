package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var i int
	i = 10000
	marshal_int, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	/* We use the json.Marshal() method to marshal and integer value.
	You will notice that the function returns a unint8 (an 8-bit unsigned integer) or byte. We need to cast it to a string */
	fmt.Println("the value of marshal_int is ", string(marshal_int))

	//You can also marshal a Boolean type using the marshal function.
	marshal_bool, _ := json.Marshal(false)
	fmt.Println("the vlaue of marshal_bool is ", marshal_bool)

	marshal_float, _ := json.Marshal(3.145)
	fmt.Println("the value of marshal_float is ", marshal_float)

	// outpurs are:
	// the value of marshal_int is  10000
	//the vlaue of marshal_bool is  [102 97 108 115 101]
	//the value of marshal_float is  [51 46 49 52 53]
}

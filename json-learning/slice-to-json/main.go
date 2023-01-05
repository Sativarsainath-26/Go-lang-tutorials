package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welocome to a class on conversion of slice to json")

	slice_raw := []string{"Microsoft", "IBM", "Amazon", "Google"}

	marshal_slice, err := json.Marshal(slice_raw)
	if err == nil {
		fmt.Println("the value of marshal_slice is given by", marshal_slice)
		fmt.Println("the value of marshal_slice in string is", string(marshal_slice))
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("doing testing for on-prem ")
	var s string

	fmt.Scanln(&s)
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

}

package main

import (
	"fmt"
	"strings"
)

// declaring a struct
func main() {
	fmt.Println("learning about string array")

	odd := []string{"{containerport = 8089, protocol = tcp, hostport= 8089}"}

	fmt.Println(odd)
	fmt.Println(len(odd))

	for _, v := range odd {
		fmt.Println(v)
		s := strings.TrimSpace(v)
		fmt.Println(s)
		if strings.Contains(v, "{}") {
			fmt.Println("Yes")
		}
	}

}

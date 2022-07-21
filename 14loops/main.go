package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on loops in golang")

	days := []string{"sunday", "tuesday", "wednesday", "thursday", "friday"} // initialized and declared a slice of string type
	fmt.Println(days)

	/*	for d := 0; d < len(days); d++ { //type 1 using for loop
			fmt.Println(days[d])
		}

		for i := range days {   // type 2 using for loop
			fmt.Println(days[i])
		} */

	for index, day := range days {
		fmt.Println("index is ", index, "and day is ", day)
	}

	rouguevalue := 1

	for rouguevalue < 10 {

		if rouguevalue == 2 {
			goto lco
		}
		if rouguevalue == 5 {
			rouguevalue++
			continue
		}

		fmt.Println("value is ", rouguevalue)
		rouguevalue++
	}

lco:
	fmt.Println("jumping at learncode.in")

}

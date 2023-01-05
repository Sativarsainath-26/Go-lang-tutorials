package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" Welcome to time study of Golang")

	presentTime := time.Now()
	//here I am just printing the date time with out any formatting
	fmt.Println(presentTime)
	// here I am printing the data time week with formatting
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// follow official docs and try to understand how it is working.
	// here I have given date and time.
	createdDate := time.Date(2022, time.July, 02, 21, 07, 2, 0, time.Local)
	fmt.Println(createdDate)

}

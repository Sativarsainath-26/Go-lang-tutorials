/* In the go , the arrray is used very less */

package main

import "fmt"

func main() {
	fmt.Println("welocome to a class of array ")

	var fruitlists [4]string // initializing a array

	fruitlists[0] = "apple" // declaring a array
	fruitlists[1] = "mango"
	fruitlists[3] = "peach"
	fmt.Println("Fruit list is :", fruitlists) // output is : Fruit list is : [apple mango  peach]
	// in the above output you are observing more space between mango and peach that is because i haven't initialized fruits[2]

	var veglist = [4]string{"tomato", "potato", "beans"}     //initializing and declaring a array
	fmt.Println("vegetable list is: ", veglist)              // output is : vegetable list is:  [tomato potato beans ]
	fmt.Println("lenght of the veglist is : ", len(veglist)) // output is  :  lenght of the veglist is :  4

}

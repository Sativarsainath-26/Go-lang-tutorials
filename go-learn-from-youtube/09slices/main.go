// slices are much most used in golang.

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to a class on slices in Golang ")

	var fruitlist = []string{"apple", "mango", "banana"}

	fmt.Printf("Type of fruitlist is given by :%T\n", fruitlist)
	fruitlist = append(fruitlist, "jamun", "Rasmalai")

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3]) // it is slicing from index 1 to end .

	fmt.Println(fruitlist)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	//highscores[4] = 674   output >>>>>> panic: runtime error: index out of range [4] with length 4
	highscores = append(highscores, 674, 345) // why here it is able to increase the slice length and give ouput with out throwing a error.
	fmt.Println(highscores)
	// when we are using append it increases the memory size that is the reason it is not throwing error

	sort.Ints(highscores) //sorts a slice in ascending order
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores)) //it returns a boolean value based on whether slice is sorted or not.

	// how to remove a value from a slice based on index.

	var courses = []string{"javascript", " react js", " python", "java"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}

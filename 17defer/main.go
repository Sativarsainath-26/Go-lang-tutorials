package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on defer keyword in golang")

	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("hello")
	Mydefer()

}

// output :
// hello,4,3,2,1,0,two,one,world
//  0, 1,2,3,4
func Mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

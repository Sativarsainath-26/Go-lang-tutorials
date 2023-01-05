package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to a class on switch-case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("number in dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you can go ahead 1 step")
	case 2:
		fmt.Println("you can go ahead 2 step")
	case 3:
		fmt.Println("you can go ahead 3 step")
	case 4:
		fmt.Println("you can go ahead 4 step")
		fallthrough //this helps if the value is 4 then it also executes the corresponding next step
	case 5:
		fmt.Println("you can go ahead 5 step")
	case 6:
		fmt.Println("you can go ahead 6 step and you can roll once again")
	default:
		fmt.Println("what was that ")
	}

}

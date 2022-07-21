package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on methods in golang")
	// whenever in the class if we are creating functions then it is referred to as methods in java .
	// In the same manner whenever you are creating functions in the struct it is referred to as methods in golang.
	// In golang there is no such concept such as classes and objects

	fmt.Println("welcome to a class on struct in golang")

	//no inheritance in golang
	//there is no super or parent class concept.

	sainath := User{"sainath", "sainath.email.com", 16, true}
	fmt.Println(sainath)
	fmt.Printf("sainath details are : %+v\n", sainath)
	fmt.Printf("Name and email is  : %v %v \n", sainath.Name, sainath.Email)

	sainath.Getstatus()
	sainath.NewEmail()

	fmt.Printf("Name and email is  : %v %v \n", sainath.Name, sainath.Email)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool // as the first letter is capital it is acessible anywhere
}

func (u User) Getstatus() { //Here we have created a method
	fmt.Println("Is user active:", u.Status)

}

func (u User) NewEmail() {
	u.Email = "sai.test.go.dev"
	fmt.Println("Emaile of the user is :", u.Email)
}

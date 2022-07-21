package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to a class on files in golang")
	content := "this is goes into a file - sainat-learnig"

	file, err := os.Create("./myfiles") // it is used for creating a file in the present directory
	/*	if err != nil {                     // it is checking if there is any error in creating a file
		panic(err)

	}*/
	checkNilErr(err)
	len, err := io.WriteString(file, content) // this function is used to write the content in to the file

	/*if err != nil {
		panic(err)
	} */
	checkNilErr(err)
	fmt.Println("length of the file is :", len)
	defer file.Close()
	ReadFile("./myfiles")

}

func ReadFile(filename string) {

	databyte, err := ioutil.ReadFile(filename) //Remember :it is reading the file from that path but it reads in the bytes not in string
	/*if err != nil {  using so many times is not a good approach so we can create a function for this and use that..
		panic(err)
	}*/
	checkNilErr(err)
	fmt.Println("Text inside the file name is ", string(databyte)) // here we are converting bytes of data into string.

}

func checkNilErr(err error) { //this is the standard way of handling errors.
	if err != nil {
		panic(err)
	}

}

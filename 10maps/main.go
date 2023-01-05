package main

import "fmt"

func main() {
	fmt.Println(" Welcome to a class on maps in golang ")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["Rb"] = "Ruby"
	languages["py"] = "python"

	fmt.Println("list of all languages : ", languages)
	fmt.Println("Js short for : ", languages["js"])

	//delete a keyword

	delete(languages, "Rb")
	fmt.Println(languages)

	// loops are interesting in golang

	for key, values := range languages {
		fmt.Printf("key is %v and value is %v\n", key, values)
	}
}

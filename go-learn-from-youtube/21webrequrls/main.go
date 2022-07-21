package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to class on get request creating in golang")
	//PerformGetRequest()
	//	PerformPostRequest()
	PerformFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)
	// this is the first way of reading the content from response body we can even define second way  that is using string builder
	//in this code it self
	/* content, _ := ioutil.ReadAll(response.Body)

	//content is in the byte format that is the reason we are doing string(content)
	fmt.Println(string(content)) */

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	bytecount, _ := responseString.Write(content)

	fmt.Println("Byte count is ", bytecount)
	fmt.Println(responseString.String())

}

func PerformPostRequest() {

	const myurl = "http://localhost:1111/post"

	//fake json payload

	requestBody := strings.NewReader(`
	    {
			"course name":"Golang programming language",
			"price":0,
			"platform":"learncode.in"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformFormRequest() {

	const myurl = "http://localhost:1111/postform"

	//form data

	data := url.Values{}
	data.Add("firstname", "sainath")
	data.Add("lastname", "sativar")
	data.Add("email", "sai.gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

package main

import (
	"fmt"
)

func main() {

	var url string
	var method string
	var option string
	var body string
	var header string

	fmt.Println("Starting HTTP request creator")
	fmt.Println("HTTP request URL: ")
	fmt.Scan(&url)
	fmt.Println("What HTTP method do you want to use: ")
	fmt.Scan(&method)

	//* HTTP REQUEST BODY
	fmt.Println("Do you want to add body to the request? [y/n]")
	fmt.Scan(&option)

	if ynToBool(option) {
		fmt.Printf("Do you want to: \n1. Type json body,\n2. Import body from json file.\n")
		if option == "1" {
			fmt.Println("Body: ")
			fmt.Scan(&body)
		}
		if option == "2" {
			fmt.Println("Path to json file: ")
		}
	}

	//* HTTP REQUEST HEADER
	fmt.Println("Header: ")
	fmt.Scan(&header)

	//* HTTP REQUEST SENDING
	fmt.Println("Do you want to send request? [y/n]")
	fmt.Scan(&option)
	if ynToBool(option) {
		fmt.Println("Sending request")
	} else {
		fmt.Println("Do you want to save request to file? [y/n]")
		fmt.Scan(&option)
		if ynToBool(option) {
			fmt.Println("Closing CLI app")
		} else if ynToBool(option) {
			fmt.Println("Path to directory where do you want to save request: ")
			//! CREATING FILE IN PATH, SAVING REQUEST
		}
	}
}

func ynToBool(option string) bool {
	if option == "y" {
		return true
	} else {
		return false
	}
}

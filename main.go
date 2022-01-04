package main

import (
	"fmt"
	"http-request-creator/controllers"
)

func main() {
	var args []string
	var option string

	fmt.Println("Starting HTTP request creator")
	args = append(args, controllers.URLmethod()...)

	//* HTTP REQUEST BODY
	fmt.Println("Do you want to add body to the request? [y/n]")
	fmt.Scan(&option)
	if ynToBool(option) {
		args = append(args, controllers.GetBody(option)...)
	}

	//* HTTP REQUEST HEADER
	fmt.Println("Do you want to add header to the request? [y/n]")
	fmt.Scan(&option)
	if ynToBool(option) {
		args = append(args, controllers.GetHeader()...)
	}

	//* HTTP REQUEST SENDING
	fmt.Println("Do you want to send request? [y/n]")
	fmt.Scan(&option)
	if ynToBool(option) {
		fmt.Println("Sending request")
		controllers.SendRequest(args)
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

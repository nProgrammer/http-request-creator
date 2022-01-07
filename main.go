//! //TODO - create README.md

package main

import (
	"fmt"
	"http-request-creator/controllers"
)

func main() {
	var args []string
	var option string

	fmt.Println("Starting HTTP request creator")
	args = append(args, controllers.URLmethod()...) //* Getting from user URL and HTTP method

	//* HTTP REQUEST BODY
	fmt.Println("Do you want to add body to the request? [y/n]")
	fmt.Scan(&option)
	if ynToBool(option) {
		args = append(args, controllers.GetBody(option)...) //* Getting from user body of http request
	}

	//* HTTP REQUEST HEADER
	fmt.Println("Do you want to add header to the request? [y/n]")
	fmt.Scan(&option)
	if ynToBool(option) {
		args = append(args, controllers.GetHeader()...) //* Getting from user header of http request
	}

	//* HTTP REQUEST SENDING
	fmt.Println("Do you want to send request? [y/n]")
	fmt.Scan(&option)
	if ynToBool(option) {
		fmt.Println("Sending request")
		controllers.SendRequest(args) //* Sending request via curl
	} else {
		fmt.Println("Do you want to save request to file? [y/n]")
		fmt.Scan(&option)
		if ynToBool(option) {
			controllers.SaveRequest(args) //* Saving request to the request.txt
		} else {
			fmt.Println("Closing CLI app")
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

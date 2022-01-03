package main

import "fmt"

func main() {
	fmt.Println("Starting HTTP request creator")
	fmt.Println("HTTP request URL: ")
	fmt.Println("What HTTP method do you want to use: ")

	//* HTTP REQUEST BODY
	fmt.Println("Do you want to add body to the request? [y/n]")

	//! if body TRUE
	fmt.Printf("Do you want to: \n1. Type json body,\n2. Import body from json file.\n")
	//!--- if option 1
	fmt.Println("Body: ")
	//!--- if option 2
	fmt.Println("Path to json file: ")
	//! if body FALSE

	//* HTTP REQUEST HEADER
	fmt.Println("Header: ")

	//* HTTP REQUEST SENDING
	fmt.Println("Do you want to send request? [y/n]")
	//! if send TRUE
	fmt.Println("Sending request")
	//! if send FALSE
	fmt.Println("Do you want to save request to file? [y/n]")
	//!--- if save FALSE
	fmt.Println("Closing CLI app")
	//!--- if save TRUE
	fmt.Println("Path to directory where do you want to save request: ")
	//! CREATING FILE IN PATH, SAVING REQUEST
}

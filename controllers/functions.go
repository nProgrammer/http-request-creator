package controllers

import (
	"fmt"
	"io/ioutil"
)

func URLmethod() []string {
	var url string
	var method string
	var args []string

	fmt.Println("HTTP request URL: ")
	fmt.Scan(&url)
	args = append(args, url)
	fmt.Println("What HTTP method do you want to use: ")
	fmt.Scan(&method)
	args = append(args, "-X")
	args = append(args, url)

	return args
}

func GetBody(option string) []string {
	var args []string
	var body string
	var path string

	fmt.Printf("Do you want to: \n1. Type json body,\n2. Import body from json file.\n")
	fmt.Scan(&option)
	if option == "1" {
		fmt.Println("Body: ")
		fmt.Scan(&body)
	}
	if option == "2" {
		fmt.Println("Path to json file: ")
		fmt.Scan(&path)
		body = loadFile(path)
	}
	args = append(args, "--body")
	args = append(args, body)

	return args
}

func GetHeader() []string {
	var header string
	var args []string

	fmt.Println("Header: ")
	fmt.Scan(&header)
	args = append(args, "--header")
	args = append(args, header)

	return args
}

func loadFile(path string) string {
	b, err := ioutil.ReadFile(path)
	checkErr(err)
	return string(b)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

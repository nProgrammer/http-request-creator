package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func URLmethod() []string {
	var url string
	var method string
	var args []string

	fmt.Println("HTTP request URL: ")
	fmt.Scan(&url)
	args = append(args, url) //* Appending URL to the slice
	fmt.Println("What HTTP method do you want to use: ")
	fmt.Scan(&method)
	args = append(args, "-X") //* Appending http method flag to the slice the curl
	args = append(args, method)

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
	args = append(args, "--body") //* Appending curl body flag to the slice
	args = append(args, body)     //* Appending body to the slice

	return args
}

func GetHeader() []string {
	var header string
	var args []string

	fmt.Println("Header: ")
	fmt.Scan(&header)
	args = append(args, "--header") //* Appending curl header flag to the slice
	args = append(args, header)     //* Appending header to the slice

	return args
}

func SendRequest(args []string) string {
	cmd := exec.Command("curl", args...)
	log.Println("curl ", args)
	data, err := cmd.Output()
	checkErr(err)
	return string(data)
}

func SaveRequest(args []string) {
	var path string

	fmt.Println("Path to directory where do you want to save request: ")
	fmt.Scan(&path)
	arguments := strings.Join(args, " ")
	request := "curl " + arguments
	log.Println(request)
	requestB := []byte(request) //* Converting request from string to byte code
	file := path + "request.txt"
	err := os.WriteFile(file, requestB, 0664) //* Writing file
	checkErr(err)
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

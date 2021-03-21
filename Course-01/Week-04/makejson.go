/*
QUESTION:

Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”,
respectively. Your program should use Marshal() to create a JSON object from the map, and then your program
should print the JSON object.

*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Function to take input from user
func takeInput(message string) string {
	fmt.Println(message)

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}

func main() {
	// Initialized map
	p := make(map[string]string)

	fmt.Println("----------------------------------------------")
	p["name"] = takeInput("Enter Your Name :=>")
	p["address"] = takeInput("Enter Address :=>")

	// Marshalling map into json with some Indent
	data, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Printf("error occured while marshaling map into json: %s", err)
		os.Exit(0)
	}
	fmt.Println("----------------------------------------------")
	fmt.Printf("JSON Data : %v", string(data))
}

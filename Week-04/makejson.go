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

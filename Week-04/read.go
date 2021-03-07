/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file has
a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname
for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each
line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program
will have a slice containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	FName string `json:"fname"`
	LName string `json:"lname"`
}

var People []Name

func collectDataFromFile(fileName string) {
	var count int64 = 0
	// read file
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("[ERROR] unable to read file: ", err)
	}

	// get access to line
	fmt.Println("[INFO] Reading your file . . .")
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		// adding line to struct
		var name Name
		line := strings.Split(scanner.Text(), " ")
		name.FName = line[0]
		name.LName = line[1]

		People = append(People, name)
		count += 1
	}

	fmt.Printf("[INFO] Total line(s) readed: %v\n", count)
	fmt.Println("----------------------------------------------------------")

}

func main() {
	var fileName string

	fmt.Printf("[INPUT] Enter file name:=> ")
	_, err := fmt.Scanf("%s", &fileName)
	if err != nil {
		fmt.Println("something went wrong: ", err)
	}

	collectDataFromFile(fileName)
	fmt.Printf("[INFO]\t\t\t FILE DATA \t\t\t\n")
	fmt.Println("----------------------------------------------------------")

	for _, name := range People {
		fmt.Printf("\tfirst_name: %v, \tlast_name: %v \n", name.FName, name.LName)
	}
}

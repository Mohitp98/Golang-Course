/*
QUESTION:
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered. Truncation is the
process of removing the digits to the right of the decimal place.

*/
package main

import (
	"fmt"
	"log"
)

var inputNumber float64

func main() {
	fmt.Println("Enter a Float number:=>")
	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		log.Printf("[Error] Invalid input !")
	}

	fmt.Printf("The number you've entered is '%v'.\n", inputNumber)
	fmt.Printf("Truncated version of '%v' is '%v'.\n", inputNumber, int64(inputNumber))
}

/*
Question:
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty
integer slice of size (length) 3. During each pass through the loop, the program prompts the user to
enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any
number of integers which the user decides to enter. The program should only quit (exiting the loop)
when the user enters the character ‘X’ instead of an integer.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	ListSize int    = 3
	ExitCode string = "X"
)

func main() {
	var input string
	var initSlice = make([]int, 3, ListSize)
	fmt.Printf("Initialized slice with Len: %v & Cap: %v \n", len(initSlice), cap(initSlice))
	for {
		fmt.Println("\t-----------------------------------")
		fmt.Printf("Current slice: %v\n", initSlice)
		fmt.Printf("Enter a number to insert ('X' to exit): ")
		if c, err := fmt.Scanf("%s", &input); err != nil {
			if c == 0 {
				log.Println("[ERROR] Empty Input !!!")
				continue
			}
			log.Println(err)
			break
		}
		if input == ExitCode {
			fmt.Println("Exiting program...")
			os.Exit(0)
		}

		intValue, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("[ERROR] Invalid Input: ->'%s'\n", input)
			continue
		}
		i, found := Find(initSlice, 0)
		if !found {
			initSlice = append(initSlice, intValue)
		} else {
			initSlice[i] = intValue

		}
		sort.Ints(initSlice)
		fmt.Println(initSlice)

	}
}

//
// Find: Finds a integer inside a slice
func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

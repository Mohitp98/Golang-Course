package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please input numbers separated by space:=> ")
	scanner := bufio.NewReader(os.Stdin)
	input, _, _ := scanner.ReadLine()
	newString := strings.Split(string(input), " ")
	var values []int
	for _, s := range newString {
		n, _ := strconv.Atoi(s)
		values = append(values, n)
	}
	BubbleSort(values)
	fmt.Println(values)
}

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j+1] < array[j] {
				swapNumbers(array, j)
			}
		}
	}
}

func swapNumbers(a []int, j int) {
	var temp int
	temp = a[j]
	a[j] = a[j+1]
	a[j+1] = temp
}

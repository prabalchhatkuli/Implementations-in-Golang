package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Swap(slice []int, i int, swapped *bool) {
	if slice[i] > slice[i+1] {
		slice[i], slice[i+1] = slice[i+1], slice[i]
		*swapped = true
	}
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			Swap(slice, j, &swapped)
		}
		if !swapped {
			break
		}
	}
}

const MAX_SIZE int = 10

func main() {
	inputArr := make([]int, 0, MAX_SIZE)

	fmt.Println("Enter an array of integer: ")

	scanner := bufio.NewScanner(os.Stdin)

	for counter := 0; counter < MAX_SIZE; counter++ {
		fmt.Printf("Enter integer %d (or press X to stop): ", counter)
		if scanner.Scan() {
			if scanner.Text() == "X" {
				break
			} else {
				value, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("Invalid input")
				} else {
					inputArr = append(inputArr, value)
				}
			}
		}
	}

	fmt.Println("Entered Unsorted array:", inputArr)

	BubbleSort(inputArr)

	fmt.Println("Sorted array:", inputArr)
}

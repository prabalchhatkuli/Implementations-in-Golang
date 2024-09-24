package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputStr string
	fmt.Println("Please enter a String with letters i, a, n in order:")
	_, err := fmt.Scan(&inputStr)
	if err != nil {
		return
	}

	lowerCaseInput := strings.ToLower(inputStr)

	var found bool = false

	if lowerCaseInput[0] == 'i' && lowerCaseInput[len(lowerCaseInput)-1] == 'n' {
		for i := 1; i < len(lowerCaseInput); i++ {
			if lowerCaseInput[i] == 'a' {
				found = true
				break
			}
		}
		if found {
			fmt.Println("Found")
		} else {
			fmt.Println("Not Found")
		}
	} else {
		fmt.Println("Not Found")
	}

}

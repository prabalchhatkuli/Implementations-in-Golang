package main

import (
	"fmt"
	"sort"
	"strconv"
)

const InputLength int = 3

func main() {
	sli := make([]int, 0, InputLength)

	var inputString string
	var elemCounter int = 0

	for true {
		fmt.Printf("Enter element %d: or press 'X' to exit: ", elemCounter)
		_, err := fmt.Scan(&inputString)
		if err != nil {
			return
		}
		if inputString == "X" {
			break
		} else {
			convertedInt, err := strconv.Atoi(inputString)
			if err != nil {
				println(err)
			} else {

				sli = append(sli, convertedInt)

				sort.Ints(sli)

				fmt.Println("Sorted slice:")
				fmt.Println(sli)
			}
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

var (
	wg sync.WaitGroup
)

const PARTITIONS int = 4

func sortIntegers(index int, slice []int) {
	fmt.Printf("Goroutine %d will sort : %v\n", index, slice)
	sort.Ints(slice)
	fmt.Printf("Goroutine %d sorted : %v\n", index, slice)
	wg.Done()
}

func main() {

	wg.Add(4)

	elementList := make([]int, 0, 20)

	fmt.Println("Please enter the elements in a new line: (Press X to stop)")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			if scanner.Text() == "X" {
				break
			} else {
				value, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("Please enter a valid number")
				} else {
					elementList = append(elementList, value)
				}
			}
		}
	}

	fmt.Printf("The input list : %v\n", elementList)

	inputLength := len(elementList)

	minimumELements := inputLength / PARTITIONS
	extraElements := inputLength % PARTITIONS

	startingIndex := 0

	for i := 0; i < PARTITIONS; i++ {

		if extraElements > 0 {
			go sortIntegers(i, elementList[startingIndex:startingIndex+minimumELements+1])
			startingIndex += minimumELements + 1
			extraElements--
		} else {
			go sortIntegers(i, elementList[startingIndex:startingIndex+minimumELements])
			startingIndex += minimumELements
		}
	}

	wg.Wait()

	sort.Ints(elementList)

	fmt.Println("Sorted List is:")
	fmt.Println(elementList)

}

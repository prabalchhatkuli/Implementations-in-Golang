package main

import "fmt"

func main() {
	var flNumber float32

	fmt.Println("Please enter a floating point number:")
	_, err := fmt.Scan(&flNumber)
	if err != nil {
		return
	}

	var convertedNumber int = int(flNumber)
	fmt.Printf("The number is %d", convertedNumber)
}

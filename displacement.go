package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GenDisplaceFn(acc, initVelocity, initDisp float64) func(float64) float64 {
	return func(t float64) float64 {
		return (acc*t*t)/2 + initVelocity*t + initDisp
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var acc, initVelocity, initDisp, time float64

	fmt.Println("Please enter acceleration a:")
	if scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		} else {
			acc = value
		}
	}

	fmt.Println("Please enter initial velocity v0:")
	if scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		} else {
			initVelocity = value
		}
	}
	fmt.Println("Please enter initial displacement s0:")
	if scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		} else {
			initDisp = value
		}
	}

	fmt.Println("Please enter time: ")
	if scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		} else {
			time = value
		}
	}

	fn := GenDisplaceFn(acc, initVelocity, initDisp)

	fmt.Printf("The displacement after the entered time is %f\n", fn(time))

}

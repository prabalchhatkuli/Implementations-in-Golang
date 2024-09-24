package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Println("Please enter the name of the file to read the data from:")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		filename = scanner.Text()
	}

	listOfNames := make([]Name, 0, 5)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		scanner = bufio.NewScanner(file)
		for scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			if len(fields) == 2 {
				currName := Name{fields[0], fields[1]}
				listOfNames = append(listOfNames, currName)
			} else {
				println("File format not as expected!")
			}
		}
	}

	for index, name := range listOfNames {
		fmt.Println(strconv.Itoa(index) + ": " + "First Name: " + name.fname + ", Last Name: " + name.lname)
	}

}

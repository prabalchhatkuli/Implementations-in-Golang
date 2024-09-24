package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var name, address string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please Enter Name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Println("Please Enter Address: ")
	if scanner.Scan() {
		address = scanner.Text()
	}

	var contactMap map[string]string

	contactMap = map[string]string{
		name: address,
	}

	contactJson, err := json.Marshal(contactMap)
	if err != nil {
		println(err)
		return
	} else {
		println(string(contactJson))
	}

}

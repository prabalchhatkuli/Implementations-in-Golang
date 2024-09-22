package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal interface {
	Eat()
	Move()
	Speak()
}

type animalData struct {
	food       string
	locomotion string
	noise      string
}

type cow struct {
	name string
	data animalData
}

func (a cow) Eat() {
	fmt.Println("==>" + a.data.food)
}

func (a cow) Move() {
	fmt.Println("==>" + a.data.locomotion)
}

func (a cow) Speak() {
	fmt.Println("==>" + a.data.noise)
}

type bird struct {
	name string
	data animalData
}

func (a bird) Eat() {
	fmt.Println("==>" + a.data.food)
}

func (a bird) Move() {
	fmt.Println("==>" + a.data.locomotion)
}

func (a bird) Speak() {
	fmt.Println("==>" + a.data.noise)
}

type snake struct {
	name string
	data animalData
}

func (a snake) Eat() {
	fmt.Println("==>" + a.data.food)
}

func (a snake) Move() {
	fmt.Println("==>" + a.data.locomotion)
}

func (a snake) Speak() {
	fmt.Println("==>" + a.data.noise)
}

func executeAction(a animal, act string) {
	switch act {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Invalid query: Enter eat/move/speak")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cowData := animalData{food: "grass", locomotion: "walk", noise: "moo"}
	birdData := animalData{food: "worms", locomotion: "fly", noise: "peep"}
	snakeData := animalData{food: "mice", locomotion: "slither", noise: "hsss"}

	dataMap := make(map[string]animalData)

	dataMap["cow"] = cowData
	dataMap["bird"] = birdData
	dataMap["snake"] = snakeData

	animalList := make([]animal, 0)

	quitRequested := false

	for {
		if quitRequested {
			break
		}
		fmt.Println("Enter your request: newanimal(eg: newanimal harry cow:bird:snake)/query(query harry eat:move:speak)/quit")
		fmt.Print(">")

		if scanner.Scan() {
			inputText := scanner.Text()
			inputWords := strings.Split(inputText, " ")
			option := inputWords[0]

			switch option {
			case "newanimal":
				switch inputWords[2] {

				case "cow":
					newCow := cow{name: inputWords[1], data: cowData}
					animalList = append(animalList, newCow)

				case "bird":
					newBird := bird{name: inputWords[1], data: birdData}
					animalList = append(animalList, newBird)

				case "snake":
					newSnake := snake{name: inputWords[1], data: snakeData}
					animalList = append(animalList, newSnake)

				default:
					fmt.Println("Invalid animal type: Enter cow/bird/snake")
				}
				fmt.Println("Created it!")

			case "query":
				isFound := false
				for _, animal := range animalList {
					switch a := animal.(type) {
					case cow:
						if a.name == inputWords[1] {
							executeAction(a, inputWords[2])
							isFound = true
						}
					case bird:
						if a.name == inputWords[1] {
							executeAction(a, inputWords[2])
							isFound = true
						}
					case snake:
						if a.name == inputWords[1] {
							executeAction(a, inputWords[2])
							isFound = true
						}
					}
				}
				if !isFound {
					fmt.Printf("An animal %s was not found\n", inputWords[1])
				}

			case "quit":
				quitRequested = true
			}

		}

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println("==>" + a.food)
}

func (a Animal) Move() {
	fmt.Println("==>" + a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println("==>" + a.noise)
}

func makeAction(a Animal, act string) {
	switch act {

	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Unknown action: " + act)
	}
}

func main() {
	//setup
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Please enter name of the Animal: options 'cow', 'bird', 'snake'. " +
			"Followed by the information you would like: options 'eat', 'move', 'speak':")

		var animal string
		var action string
		fmt.Print(">")
		if scanner.Scan() {
			inputText := scanner.Text()
			inputWords := strings.Split(inputText, " ")
			if len(inputWords) == 2 {
				animal = inputWords[0]
				action = inputWords[1]

				switch animal {
				case "cow":
					makeAction(cow, action)
				case "bird":
					makeAction(bird, action)
				case "snake":
					makeAction(snake, action)
				default:
					fmt.Println("Unknown animal: " + animal)
				}

			} else {
				fmt.Println("Invalid input")
			}
		}
	}
}

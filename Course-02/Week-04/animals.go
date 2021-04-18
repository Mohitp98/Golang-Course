package main

import (
	"fmt"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (ani animal) Eat() {
	fmt.Println(ani.food)
	return
}

func (ani animal) Move() {
	fmt.Println(ani.locomotion)
	return
}

func (ani animal) Speak() {
	fmt.Println(ani.noise)
	return
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}
	var subjectAnimal animalInterface
	for {
		var command, requestAni, requestType string
		fmt.Print(">")
		_, err := fmt.Scan(&command, &requestAni, &requestType)
		if err != nil {
			fmt.Println("Error Occurred while reading input: ", err)
		}
		if command == "query" {
			subjectAnimal = animalMap[requestAni]
			switch requestType {
			case "eat":
				subjectAnimal.Eat()
			case "move":
				subjectAnimal.Move()
			case "speak":
				subjectAnimal.Speak()
			}
		}
		if command == "newanimal" {
			animalMap[requestAni] = animalMap[requestType]
			fmt.Println("Created it!")
		}
	}
}

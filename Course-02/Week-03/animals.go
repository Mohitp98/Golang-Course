package main

import (
	"fmt"
	"os"
)

type Animal struct {
	food, locomotion, sound string
}

func (v Animal) Eat() {
	fmt.Println(v.food)
}

func (v Animal) Move() {
	fmt.Println(v.locomotion)
}

func (v Animal) Speak() {
	fmt.Println(v.sound)
}

func main() {
	data := map[string]Animal{
		"bird":  Animal{"worms", "fly", "peep"},
		"cow":   Animal{"grass", "walk", "moo"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for {
		fmt.Print(">")
		subjectAnimal := "0"
		animalAction := "0"
		_, err := fmt.Scan(&subjectAnimal, &animalAction)
		if err != nil {
			fmt.Println("Invalid input : ", err)
			os.Exit(0)
		}
		if animalAction == "eat" {
			data[subjectAnimal].Eat()
		} else if animalAction == "move" {
			data[subjectAnimal].Move()
		} else if animalAction == "speak" {
			data[subjectAnimal].Speak()
		}
	}
}

package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Print(a.food)
}

func (a Animal) Move() {
	fmt.Print(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Print(a.noise)
}

func main() {
	var animalName string
	var animalAction string
	var animal Animal

	fmt.Println("Animals: cow, bird, snake \n Actions: eat, move, speak. \n Please, enter animal and action, like: > cow eat")

	for {
		fmt.Print("> ")
		fmt.Scan(&animalName, &animalAction)
		if animalName == "cow" {
		
			animal = Animal{food: "grass", locomotion: "walk", noise: "moo"}
		}
		if animalName == "bird"{
			animal = Animal{food: "worms", locomotion: "fly", noise: "peep"}
		}
		if animalName ==  "snake"{
			animal = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
		}


		if animalAction == "eat"{
			animal.Eat()
		}
		if animalAction == "move"{
			animal.Move()
		}
		if animalAction == "speak"{
			animal.Speak()
		}
	}

}

package main

import "fmt"

type Sounder interface {
	makeSound()
}

type Dog struct {
	sound string
}
func (d Dog) makeSound() {
	fmt.Println(d.sound)
}

type Duck struct {
	sound string
}
func (d Duck) makeSound() {
	fmt.Println(d.sound)
}

func doTheTrick(animal Sounder) {
	animal.makeSound()
}

func main() {
	duck := Duck{sound: "Quack"}
	dog := Dog{sound: "Bark"}

	doTheTrick(duck)
	doTheTrick(dog)
}
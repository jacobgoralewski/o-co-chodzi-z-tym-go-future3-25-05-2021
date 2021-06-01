package main

import "fmt"

type Cat struct {
	sound string
}

func (c Cat) makeSound() {
	fmt.Println(c.sound)
}

func main() {
	englishCat := Cat{sound: "Meow"}
	englishCat.makeSound()

	polishCat := Cat{sound: "Miał"}
	polishCat.makeSound()
}
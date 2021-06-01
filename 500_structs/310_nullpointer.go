package main

import "fmt"

type Bird struct {
	sound string
}

func (b *Bird) makeSound() {
	fmt.Println("Tweet")
}

func main() {
	b := &Bird{}
	b = nil

	b.makeSound()
}
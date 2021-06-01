package main

import "fmt"

type Person struct {
	name string
}

func (p Person) String() string {
	return fmt.Sprintf("Person<name: %s>", p.name)
}

func main() {
	p := Person{name: "Jakub"}
	fmt.Println(p)

}
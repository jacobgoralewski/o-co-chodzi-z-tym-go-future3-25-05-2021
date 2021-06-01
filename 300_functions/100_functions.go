package main

import "fmt"

func Greet(name string) {
	fmt.Printf("Hello %s\n", name)
}

func main() {
	var name = "Mark"
	Greet(name)
}

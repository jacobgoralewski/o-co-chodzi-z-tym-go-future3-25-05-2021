package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	p := Person{
		Name:    "Jacob",
		Age:     26,
		Address: "Gdańsk",
	}

	serializedPerson, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Json: %s", string(serializedPerson))
}

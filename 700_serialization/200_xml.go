package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Employee struct {
	Name   string `xml:"name"`
	Age    int    `xml:"age"`
	Salary int    `xml:"salary"`
}

func main() {
	e := Employee{
		Name:    "Jacob",
		Age:     26,
		Salary: 15000,
	}

	serializedEmployee, err := xml.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Json: %s", string(serializedEmployee))
}

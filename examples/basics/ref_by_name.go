package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int32
}

func main() {
	p := Person{"Viktor", 23}
	p2 := Person{age: 21, name: "Zsu"}
	fmt.Println("Person:", p)
	fmt.Println("P2:", p2)

	// create with default values
	p3 := new(Person)
	fmt.Println("P3:", p3)
}

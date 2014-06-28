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
	p2 := Person{age: 22, name: "Zsu"}
	fmt.Println("Person: %v", p)
	fmt.Println("P2: %v", p2)
}

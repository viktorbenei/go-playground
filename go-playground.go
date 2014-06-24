package main

import (
	"fmt"
	"github.com/viktorbenei/go-playground/mclass"
)

func main() {
	fmt.Printf("Hello Playground!\n")
	person := mclass.Person{Name: "Viktor"}
	fmt.Printf("Name: %v\n", person.GetName())
}

/*
Reference by name

A simple struct test, creating objects from structs with either positional parameters or through named parameters.
*/

package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int32
}

//
// Base Println doesn't support the formatting Printf does
//  but adding \n to every Printf call is not DRY (and not fun at all)
// So this function does exactly this: you can use
//  exactly the same way as you can Printf but it will
//  automatically add a newline to the end of the line
//
func Printfln(format string, a ...interface{}) (n int, err error) {
	format = format + "\n"
	return fmt.Printf(format, a...)
}

func main() {
	p := Person{"Viktor", 23}
	p2 := Person{age: 21, name: "Zsu"}
	Printfln("Person: %#v", p)
	Printfln("P2: %#v", p2)

	// create with default values
	p3 := new(Person)
	Printfln("P3: %v", p3)
}

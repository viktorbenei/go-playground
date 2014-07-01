/*
Logging examples

Simple logging examples
*/

package main

import (
	"fmt"
	"log"
)

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
	defer Printfln("-> This should run last, but should run regardless. However it won't run if the program exits - as log.Fatal does.")
	fmt.Println("Start")
	Printfln("My format thing %d", 12)
	log.Fatal("This will exit right away")
	fmt.Println("This won't be logged")
}

/*
Reads a file line-by-line and transforms it
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	filePath    = flag.String("filepath", "", "[REQUIRED] input file's path")
	prefixThis  = flag.String("prefix", "", "prefix with this")
	postfixThis = flag.String("postfix", "", "postfix with this")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [FLAGS]", os.Args[0])
	flag.PrintDefaults()
}

func transformLine(line, pref, postf string) string {
	return pref + line + postf
}

func main() {
	//	var err error

	flag.Usage = usage
	flag.Parse()

	if *filePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	//fmt.Println(*filePath, postfixThis, prefixThis)

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal("Could not open the input file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = transformLine(line, *prefixThis, *postfixThis)
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

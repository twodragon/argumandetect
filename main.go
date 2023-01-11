package main

import (
	"fmt"
	"os"
)

func main() {

	// The first argument

	// is always program name

	myProgramName := os.Args[0]
	toGetAllArgs := os.Args[1:]

	fmt.Println(myProgramName)
	fmt.Println(toGetAllArgs)
	
	var first string
	// Taking input from user
	fmt.Scanln(&first)
	fmt.Println("x: ")
}

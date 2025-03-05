package main

import (
	"fmt"
	"math"
	"strings"
)


func main() {
	fmt.Println(math.Floor(3.12321312))
	fmt.Println(strings.Title("How do strings work??"))
	// Lets see how the runes behave
	fmt.Println("String Literals")
	fmt.Println("A")  // this is a string?
	fmt.Println("\n") // This also is a string?
	fmt.Println("Runes are displayed bellow")
	fmt.Println('A')
	fmt.Println('\n')
	fmt.Println('x')
}

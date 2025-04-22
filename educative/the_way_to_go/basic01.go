package main

import (
	"fmt"
)

func main() {
	var message string = "arun dev"
	fmt.Printf("The message is :: %s", message)
	fmt.Printf("the address is : %p", &message)
	fmt.Printf("the value at the address is : %s", *&message)
}

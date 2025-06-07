package greetings

import "fmt"

func ExcitedGreet() {
	fmt.Println("Woohooo!! Welcome to this.")
}

func SuperGreet() {
	var name string = "Stan"
	fmt.Printf(" Hey there : %s\n", name)
	var namePtr *string = &name
	fmt.Printf("\t The address of where the name is at is :: {%p}\n", namePtr)
	// This is pointer address listing and the dereferencing of the pointer
	fmt.Printf("\t The value stored at address {%p} is {%s}\n", namePtr, *namePtr)
}

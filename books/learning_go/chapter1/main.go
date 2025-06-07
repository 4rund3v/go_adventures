package main

import (
	"chapter1/helloWorld"
	"fmt"
	"greetings"
	"numerics"
)

func main() {
	helloWorld.Greet()
	greetings.ExcitedGreet()
	greetings.SuperGreet()

	var equiTriangle = numerics.NewTriangle(6, 6, 6)
	fmt.Printf("The equilateral triangle initialized is :: {%v}\n", equiTriangle)
	equiTriangle.CalculatePerimeter()
	equiTriangle.CalculateArea()
}

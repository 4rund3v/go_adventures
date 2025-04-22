package main

import (
	"fmt"
)

// Global variable
var total int

func mainAdd() {
	// local variable
	var sum int
	{
		// Block level local variable
		count := 5
		sum = calculateSum(&count, 2)
		fmt.Println("Sum:", sum)
	}
	total = sum
	fmt.Printf("post block 1 Sum: %d Total :%d\n\n", sum, total)
	{
		count := 10
		sum = calculateSum(&count, 3)
		fmt.Println("Sum:", sum)
	}
	fmt.Printf("post block 2 Sum: %d Total :%d\n\n", sum, total)
	total := total + sum
	fmt.Println("Total:", total)
}

func calculateSum(count *int, factor int) int {
	// Accessing local variable
	sum := *count * factor
	return sum
}
func main() {
	var message string = "arun dev"
	const greeting string = "Hey"
	const exclaim rune = '!'
	const (
		name   string = "arun"
		suffix rune   = '!'
	)
	fmt.Println(name, suffix)
	fmt.Printf("%s%c\n", name, suffix)

	fmt.Printf("The message is :: %s", message)
	fmt.Printf("the address is : %p", &message)
	fmt.Printf("the address of greeting is : %s", greeting+message+string(exclaim))

	fmt.Printf("the value at the address is : %s", *&message)
	const specialValue1 = iota
	fmt.Println("the special value is:: ", specialValue1)
	const specialValue2 = iota
	fmt.Println("the special value 2 is:: ", specialValue2)
	const specialValue3 = iota
	fmt.Println("the special value 3 is:: ", specialValue3)
	const (
		specialValueParenthesised1 = iota
		specialValueParenthesised2 = iota
		specialValueParenthesised3 = iota
	)
	fmt.Println("the special values is:: ", specialValueParenthesised1, specialValueParenthesised2, specialValueParenthesised3)
	type Age int8
	// fmt.Println("value of age is :: ", Age)
	const arunAge Age = 30
	fmt.Printf("age of arun is : %d", arunAge)
	const (
		specialValueParenthesisedV1 Age = iota
		specialValueParenthesisedV2
		specialValueParenthesisedV3
	)
	fmt.Println("the special values v1 is:: ", specialValueParenthesisedV1,
		specialValueParenthesisedV2, specialValueParenthesisedV3)

	var choice bool
	var wins int
	var total float32
	fmt.Println(choice, wins, total)
	mainAdd()
	var a = float32((100 * 9.0 / 5.0) + 32)
	fmt.Println("the value of a", a)

}

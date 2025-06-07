package numerics

import (
	"fmt"
	"math"
)

type Triangle struct {
	length  int
	breadth int
	height  int
}

func NewTriangle(length, breadth, height int) *Triangle {
	var newTriangle = Triangle{length: length, breadth: breadth, height: height}
	return &newTriangle
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle(length: %d, breadth: %d, height: %d)", t.length, t.breadth, t.height)
}

func (t Triangle) CalculatePerimeter() int {
	fmt.Println("Calculating the perimeter for the triangle ::")
	var perimeter int = t.height + t.breadth + t.length
	fmt.Printf("\t Calculated the perimeter as :: {%d}\n", perimeter)
	return perimeter
}

func (t Triangle) claculateSemiPerimeter() float32 {
	var semiPerimeter float32 = float32(t.length+t.breadth+t.height) / 2
	fmt.Printf("\t The semiperimeter is :: {%f}\n", semiPerimeter)
	return semiPerimeter
}

func (t Triangle) CalculateArea() float32 {
	var area float32
	var semiPerimeter float32 = t.claculateSemiPerimeter()
	area = semiPerimeter * (semiPerimeter - float32(t.length)) * (semiPerimeter * float32(t.breadth)) * (semiPerimeter * float32(t.height))
	area = float32(math.Sqrt(float64(area)))
	fmt.Printf("\t The area calculated is : {%f}\n", area)
	return area
}

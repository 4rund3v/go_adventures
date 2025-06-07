package shapes

import (
	"fmt"
)

func PrintShapes() {
	var supportedShapes = []*Triangle{NewTriangle(0, 0, 0)}
	fmt.Printf("the list of shapes supported are :: {%v} \n", supportedShapes)
}

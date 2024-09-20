package examplepackage

import "fmt"

func IfCondition() {
	x, y := -1, false
	//x condition
	if x > 0 {
		fmt.Println("x is greater than 0")
	} else {
		fmt.Println("y is smaller than 0")
	}
	//y condition
	if y {
		fmt.Println("y is true")
	} else {
		fmt.Println("y is false")
	}
}

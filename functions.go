package main

import (
	"fmt"
)

//A simple function
func simpleFunction() {
	fmt.Println("hello world")
}
//A function with Params
func functionWithParameters(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}
//A function with Types return
func functionWithTypes(x float32, y float32) float32 {
	var total float32
	total = x + y 
	return total
}
//A function with a return named value
func functionWithNamedValue(x float32, y float32) (area float32)  {
	area = x * y
	return area
}
//Dereferencing function
func DereferenceFunction(x *int, y *string){
	*x = *x + 5
	*y = *y + " check"
	return
}
func BasicFunction() {
	// simpleFunction()
	// functionWithParameters(2, 3)
	// fmt.Println(functionWithTypes(2.0, 9.5))
	
	var age int = 3
	var name string = "Tri"
	fmt.Println("Before: ", age, name)
	DereferenceFunction(&age, &name)
	fmt.Println("After: ", age, name)

}
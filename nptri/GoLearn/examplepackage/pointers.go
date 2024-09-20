package examplepackage

import (
	"fmt"
)

func pointerProg() {
	//Declare a pointer
	var PointerTo *int

	//Declare a value
	valueToPoint := 2

	//Reference to a value
	PointerTo = &valueToPoint

	//This print an address, output different everytime
	fmt.Println(PointerTo)

	//Dereference - output will be value
	fmt.Println(*PointerTo)

	//~~~~~~~~~~//
	//Initiate a new Pointer using "new" keyword
	NewPointer := new(int)
	*NewPointer = 12052001
	fmt.Println(*NewPointer)
	fmt.Println(NewPointer)
}

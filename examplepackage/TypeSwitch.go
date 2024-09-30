package examplepackage

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2) //%v value in default format
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v)) //%q single-quote character 
	default:
		fmt.Printf("I don't know about type %T!\n", v) //%T is type bool pass in
	}
}

func TypeSwitch() {
	do(21)
	do("hello")
	do(true)
}
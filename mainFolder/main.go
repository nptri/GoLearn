package main

import cust "example.com/custompackage"
import custFunc "example.com/examplepackage"
import arraysDemoLearn "example.com/arrays"
import "fmt"

func main() {
	fmt.Println(cust.AddFunc(1, 3))
	custFunc.ForLoop()
	custFunc.TypeInference()
	arraysDemoLearn.ArrayLearn()
}

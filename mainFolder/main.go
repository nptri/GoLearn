package main

import cust "example.com/custompackage"
import custFunc "example.com/examplepackage"
import "fmt"

func main() {
	fmt.Println(cust.AddFunc(1, 3))
	custFunc.ForLoop()
	
}

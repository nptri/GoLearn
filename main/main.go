package main

import "fmt"
import cust "github.com/nptri/GoLearn/example"
import custFunc "github.com/nptri/GoLearn/custompackage"

func main() {
	fmt.Println(custFunc.AddFunc(2,3))
	cust.ForLoop()
}

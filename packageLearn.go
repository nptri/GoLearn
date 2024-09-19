package main

import (
	"fmt"
	"strings"
)
import cust "src/custompackage"

// import cust "src/custompackage" //set cust as namespace for custom import


func main() {
	lower := strings.ToLower("GOLANG GO")
	fmt.Println(lower)
	upper := strings.ToUpper("golang strings")
	fmt.Println(upper)

	AddReturnValue := cust.AddFunc(2, 3)
	fmt.Println(AddReturnValue)
}

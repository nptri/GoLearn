package example

import (
	"fmt"
	"strings"
)

// import cust "src/custompackage" //set cust as namespace for custom import


func packageLearn() {
	lower := strings.ToLower("GOLANG GO")
	fmt.Println(lower)
	upper := strings.ToUpper("golang strings")
	fmt.Println(upper)
}

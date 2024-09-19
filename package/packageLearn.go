package main
import (
	"fmt"
	"strings"
	"custompackage"
)


func main() {
	lower := strings.ToLower("GOLANG GO")
	fmt.Println(lower)
	upper := strings.ToUpper("golang strings")
	fmt.Println(upper)

	AddReturnValue := AddFunc(2, 3)
	fmt.Println(AddReturnValue)
}

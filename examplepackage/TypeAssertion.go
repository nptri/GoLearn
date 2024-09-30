package examplepackage
import "fmt"

func TypeAssertion() {
	var i interface{} = "hello String"

	s := i.(string)
	fmt.Println(s)

	s, boolVal := i.(string)
	fmt.Println(s, boolVal)

	f, boolVal := i.(float64)
	fmt.Println(f, boolVal)

	x := i.(float64) //panic and exit
	fmt.Println(x)
}
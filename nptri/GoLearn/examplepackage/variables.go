package examplepackage

import (
	"fmt"
	"math"
)

const s string = "constant"

func variablesName() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(s)

	const n = 500000000
	const z = 3e20 / n
	fmt.Println(z)
	fmt.Println(int64(z))
	fmt.Println(math.Sin(n))
}

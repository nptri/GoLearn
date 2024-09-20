package examplepackage
import "fmt"
//Anonymous function with parameters
// var (
// 	area = func(length int, width int) int {
// 		return length * width
// 	}
// )

func advanceFunctions() {
	// fmt.Println("area is:", area(5, 10))
	fmt.Printf("Area is: %.1f",
		func(coorX float32, coorY float32) (area float32) {
			area = coorX * coorY
			return
		}(20, 30),
	)
}

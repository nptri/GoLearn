package arrays
import "fmt"

func ArrayLearn () {
	var arraysDemo [4]string //we defined an array named "arraysDemo" with size 4 with string type
	arraysDemo[0] = "Hello"
	arraysDemo[1] = "World"
	arraysDemo[2] = "and"
	arraysDemo[3] = "TriNguyen"
	for i := 0; i < 4; i++ {
		fmt.Println(arraysDemo[i], " ")
	}
}
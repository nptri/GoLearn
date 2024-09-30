package arrays

import "fmt"

func GoMaps() {
	var studentNumber = map[string]int{ //a map with param is string, with value of int type
		"Anh": 1,
		"Cong": 2,
		"Ha": 3,
	}
	fmt.Println(studentNumber)
	fmt.Println("Before num:", studentNumber["Anh"])
	delete(studentNumber, "Anh")
	fmt.Println(studentNumber)
	fmt.Println("After num:", studentNumber["Anh"])

}


package examplepackage

import (
	"fmt"
)

func switchCase() {
	// var today = time.Now()
	var today = 20
	switch today {
		case 17:
			fmt.Println("Today is 17, go to work")
		case 20:
			fmt.Println("Salary day")
	}
}

package main
import "fmt"
//Anonymous function with parameters
// var (
// 	area = func(length int, width int) int {
// 		return length * width
// 	}
// )

func ForLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Println("i =", i)
	}
	//some patterns we usually see - Three component for loop
	var sum int = 0
	for y := 0; y <= 10; y++ {
		sum += 1
	}
	fmt.Println(sum)

	//While loop - it's a for loop but cut out init and post statements - if n is smaller than the conditions
	//then it is running, until the threshold then stop, if n is bigger, finish the loop
	n := 2
	for n < 5 {
		n += 1
		fmt.Println("n is:", n)
	}
	//fail case
	m := 3
	for m < 2 {
		m +=1
		fmt.Println("n is:", m) //will not print
	}
}

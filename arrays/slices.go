package arrays

import "fmt"

func Slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //slices is a way to access to a range in an array
	//define as "var name []type = array[low:high]"
	//it will slice before to low, and high to after
	fmt.Println(s)
}
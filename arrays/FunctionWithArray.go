package arrays

func FunctionWithArray(ArrayParseIn *[2]string) (FullString string) {
	for i := 0; i < 1; i++ {
		FullString += ArrayParseIn[i] 
	}
	return FullString
}
//we can use FullString = append(FullString, "new value to be added in")
//it will create a new FullString array that has "new value to be added in"


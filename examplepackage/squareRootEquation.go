package examplepackage

import "fmt"
import "errors"
type ErrorHandling struct {
	message_variable string

	num float64
}

func (errorName *ErrorHandling) Error() string {
	return errorName.message_variable
}
func denyLetter(letter string) (bool, error){
	if letter == "a" {
		return false, &ErrorHandling{
			message_variable: fmt.Sprintln("letter a is not allowed!"),
			num: 0,
		}
	} 
	return true, nil
}

func TestError() {
	result, denyError := denyLetter("a")
	if denyError != nil {
		var InternalErr *ErrorHandling
		if errors.As(denyError, &InternalErr){
			fmt.Printf("%s", InternalErr.message_variable)
		} else {
			fmt.Printf("A weird error occured")
		}
		return 
	}
	fmt.Printf("%v", result)
}
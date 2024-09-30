package examplepackage

import ("fmt"
	"errors"
)

type UnsupportedError struct {
	num1 int
	num2 int
	MessageError string
}
func (e *UnsupportedError) Error() string {
	return e.MessageError
}

func Minus(a, b int) (int, error) {
	if b < 0 {
		return 0, &UnsupportedError{
			MessageError: fmt.Sprintf("We are not supporting negative number. This case '%d'", b),
			num1: a,
			num2: b,
		}
	}
	return a + b, nil
}

func ErrorMinus() {
	a := 0
	b := -1
	result, errorDetails := Minus(a, b)
	if errorDetails != nil {
		var MinusErr *UnsupportedError
		if errors.As(errorDetails, &MinusErr){
			fmt.Printf("%s", MinusErr.MessageError)
		} else {
			fmt.Printf("unexpected error happened %s\n", MinusErr.MessageError)
		}
		return
	}
	fmt.Printf("result is: %d", result)
	
}
package errors1

import "fmt"
import "errors"

func fun1(age int) (int, error) {
	if age == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return age + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func fun2(age int) (int, error) {
	if age == 42 {
		return -1, argError{age, "Can't work with it"}
	}
	return age + 3, nil
}

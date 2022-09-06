package errorhandling

import (
	"fmt"
	"log"
)

type customError[T any] struct {
	arg     T
	problem string
}

func (e *customError[T]) Error() string {
	return fmt.Sprintf("arg: %#v, problem: %v", e.arg, e.problem)
}

// driver code
func f1(arg int) (int, error) {
	if arg <= 0 {
		return -1, &customError[int]{arg: arg, problem: "Non-positive numbers are not supported"}
	}
	return arg, nil
}

func f2(arg string) (string, error) {
	if arg == "" {
		return "", &customError[string]{arg: arg, problem: "Empty strings are not supported"}
	}
	return arg, nil
}

func Run() {

	_, err := f1(-5)
	if err != nil {
		log.Printf("f1 resulted in error: %v", err)
	}

	_, err = f2("")
	if err != nil {
		log.Printf("f1 resulted in error: %v", err)
	}

	if res, err := f1(-5); err != nil {
		log.Printf("f1 resulted in error: %v", err)
	} else {
		log.Printf("f1 result: %d", res)
	}

	_, err = f1(-2)

	if ce, ok := err.(*customError[int]); ok {
		log.Printf("f1 with arg: %v, resulted in error with problem: %v", ce.arg, ce.problem)
	} else {
		log.Printf("%v, %v, %v", ce, ok, err)
	}

}

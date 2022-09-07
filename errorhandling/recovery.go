package errorhandling

import (
	"errors"
	"fmt"
)

// recover must be called from a deferred function.
// if the outer (enclosing) function panics, the deferred one will activate
// and a `recover` call in it will catch the panic

func problematicFunc() {
	panic(errors.New("some problem"))
}

func RunRecoveryExample() {

	fmt.Println("before panic")

	defer func() {
		fmt.Println("deferred msg")

		if r := recover(); r != nil {
			fmt.Printf("recovered from panic r: %#v\n", r)
		}

		fmt.Println("after recover")
	}()
	problematicFunc()

	// this line will not execute
	fmt.Println("after panic")

}

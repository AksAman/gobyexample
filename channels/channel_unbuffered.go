package channels

import (
	"fmt"
)

func RunUnbuffered() {
	c := make(chan string)

	go func() {
		c <- "some message"
	}()

	msg, more := <-c
	fmt.Printf("msg: %v, more: %v\n", msg, more)
}

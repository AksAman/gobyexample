package timemachine

import (
	"fmt"
	"time"
)

func RunTimeOutExample() {
	c1 := make(chan string)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result-1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result-2"
	}()

	select {
	case m := <-c1:
		fmt.Printf("m: %v\n", m)
	case <-time.After(time.Second * 1):
		fmt.Printf("time limit exceeded 1\n")
	}

	select {
	case <-time.After(time.Second * 3):
		fmt.Printf("time limit exceeded 2\n")
	case m := <-c2:
		fmt.Printf("m: %v\n", m)
	}
}

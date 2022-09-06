package channels

import (
	"fmt"
	"time"
)

func task(i int) int {
	time.Sleep(time.Second)
	return i * i
}

func RunRangeChannels() {
	c := make(chan int)

	go func(n int) {
		for i := 0; i <= n; i++ {
			result := task(i)
			c <- result
		}
		close(c)
	}(5)

	for elem := range c {
		fmt.Printf("elem: %v\n", elem)
	}
}

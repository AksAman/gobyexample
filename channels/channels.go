package channels

import (
	"fmt"
)

func Run() {

	messagesChan := make(chan string)

	go func(msg string) {
		for i := 0; i < 3; i++ {
			messagesChan <- fmt.Sprintf("to channel %v: %d", msg, i)
		}
	}("ping")

	fmt.Println(<-messagesChan)
	fmt.Println(<-messagesChan)
	fmt.Println(<-messagesChan)
}

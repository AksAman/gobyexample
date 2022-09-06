package channels

import (
	"fmt"
	"time"
)

func fastWorker(done *chan string) {
	time.Sleep(time.Second * 1)
	*done <- "fast done"
}

func slowWorker(done chan string) {
	time.Sleep(time.Second * 3)
	done <- "slow done"
}

func RunChannelSelect() {

	fastChan := make(chan string)
	slowChan := make(chan string)

	go fastWorker(&fastChan)
	go slowWorker(slowChan)

	for i := 0; i < 2; i++ {

		select {
		case m2 := <-slowChan:
			fmt.Println(m2, i)
		case m := <-fastChan:
			fmt.Println(m, i)
		}
	}

}

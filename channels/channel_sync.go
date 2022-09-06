package channels

import (
	"fmt"
	"time"
)

func worker(done chan<- bool) {
	fmt.Println("Working")
	time.Sleep(time.Second * 3)
	fmt.Println("Work done")
	done <- true
}

func RunChannelSynced() {

	doneChannel := make(chan bool)

	go worker(doneChannel)
	fmt.Println("work started")
	fmt.Println("work started")
	fmt.Println("work started")
	fmt.Println("work started")

	finished := <-doneChannel
	fmt.Println(finished)
	fmt.Println("DONE")

}

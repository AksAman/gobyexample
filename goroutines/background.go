package goroutines

import (
	"fmt"
	"os"
	"time"
)

func backgroundWorker(done chan<- bool) {
	time.Sleep(time.Second * 2)

	if f, err := os.Create("./testfile"); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("f: %v\n", f)
		f.Close()
	}
	done <- true
}

func RunInBackground() {

	messageChan := make(chan bool)

	go backgroundWorker(messageChan)

	// <-messageChan
	// time.Sleep(time.Second * 4)
}

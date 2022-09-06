package timemachine

import (
	"fmt"
	"time"
)

// using timers over sleep gives
// 1. can be used as a channel across goroutines
// 2. can be cancelled

func RunTimers() {
	timer := time.NewTimer(time.Second / 2)

	<-timer.C
	fmt.Printf("Timer 1 fired\n")

	doneChan := make(chan bool)
	timer = time.NewTimer(time.Second * 5)
	stopTimer := time.NewTimer(time.Second * 1)

	go func() {
		select {
		case <-timer.C:
			{
				fmt.Printf("Timer 2 fired\n")
				// doneChan <- true
			}
		case <-stopTimer.C:
			{
				stopped := timer.Stop()
				fmt.Printf("Stop Timer fired, timer 2 stopped: %v\n", stopped)
				// doneChan <- true
			}
		}
		doneChan <- true
	}()

	<-doneChan
}

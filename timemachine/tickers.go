package timemachine

import (
	"fmt"
	"time"
)

/**
doneChan is sent 'true' on a different goroutine because
it is an unbuffered channel, hence send and receive should happen
simultaneously or it will block the for loop
**/

func RunTickers() {
	tickInterval := time.Millisecond * 500
	totalInterval := tickInterval * 5
	doneChan := make(chan bool)
	finalDoneChan := make(chan bool)

	ticker := time.NewTicker(tickInterval)
	programTimer := time.NewTimer(totalInterval)
	i := 0

	go func() {
		for {
			select {
			case <-doneChan:
				fmt.Println("\tProgram done")
				finalDoneChan <- true
				return
			case <-ticker.C:
				fmt.Println("\ttick", i)
			case <-programTimer.C:
				fmt.Println("\ttotalTimer Done")
				ticker.Stop()
				go func() {
					doneChan <- true
				}()
			}

			fmt.Println("for loop: ", i)
			i++
		}
	}()

	<-finalDoneChan

	fmt.Println("Ticker stopped")

}

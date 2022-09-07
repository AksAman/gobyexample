package channels

import (
	"fmt"
	"time"
)

func jobworker(jobs *chan string, done *chan bool) {
	for {
		// receives jobs from jobs channel
		job, more := <-*jobs
		if more {
			time.Sleep(time.Second)
			fmt.Printf("job rcvd: %v\n", job)
		} else {
			*done <- true
			fmt.Printf("All jobs rcvd, %v, %v\n", job == "", more)
			return
		}
	}
}

func RunCloseChannels() {
	jobsChannel := make(chan string)
	doneChannel := make(chan bool)

	go jobworker(&jobsChannel, &doneChannel)

	jobsCount := 3

	for i := 0; i < jobsCount; i++ {
		job := fmt.Sprintf("JOB: %d", i)
		jobsChannel <- job
		fmt.Println("sent job", job)
	}
	close(jobsChannel)
	fmt.Println("sent all jobs")

	finished := <-doneChannel
	fmt.Printf("main: jobs finished: %v\n", finished)
}

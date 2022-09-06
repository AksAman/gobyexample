package workerpools

import (
	"fmt"
	"time"
)

type job int
type jobResult struct {
	j job
	r int
}

func someExpensiveTask(j job) jobResult {
	time.Sleep(time.Second * 1)
	return jobResult{j: j, r: int(j * j)}
}

func worker(wid int, jobsChan <-chan job, resultsChan chan<- jobResult) {
	for j := range jobsChan {
		fmt.Println("\tworker", wid, "started job:", j)
		result := someExpensiveTask(j)
		fmt.Println("\tworker", wid, "finished job:", j)
		resultsChan <- result
	}
}

func RunWorkers() {
	numJobs := 10
	numWorkers := 5

	bufferSize := 0
	if numJobs > numWorkers {
		bufferSize = numJobs - numWorkers
	}
	jobsChannel := make(chan job, bufferSize)
	resultsChannel := make(chan jobResult)

	// create workers
	for i := 0; i < numWorkers; i++ {
		go worker(i, jobsChannel, resultsChannel)
	}

	// publish jobs
	for i := 0; i < numJobs; i++ {
		jobsChannel <- job(i)
	}
	close(jobsChannel)

	// receive jobResults
	results := make(map[job]jobResult)
	for i := 0; i < numJobs; i++ {
		jr := <-resultsChannel
		results[jr.j] = jr
	}

	fmt.Printf("results: %v\n", results)

}

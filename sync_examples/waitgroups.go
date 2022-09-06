package sync_examples

import (
	"fmt"
	"sync"
	"time"
)

func worker(wid int, resultsChan chan int) {
	fmt.Println("worker starting: ", wid)
	time.Sleep(time.Second * 1 / 2)
	if resultsChan != nil {
		resultsChan <- wid * wid
	}
	fmt.Println("worker done: ", wid)
}

func RunWithoutWaitGroups() {
	workersCount := 5

	resultsChan := make(chan int)

	for i := 0; i < workersCount; i++ {
		go worker(i, resultsChan)
	}

	for i := 0; i < workersCount; i++ {
		m := <-resultsChan
		fmt.Printf("result m: %v\n", m)
	}
	close(resultsChan)

	_, more := <-resultsChan
	if more {
		fmt.Println("more more")
	} else {
		fmt.Println("no more")

	}
}

func RunWithWaitGroups() {
	workersCount := 5
	var wg sync.WaitGroup

	for i := 0; i < workersCount; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i, nil)
		}()
	}
	wg.Wait()
}

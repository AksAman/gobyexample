/**
in the mutex or atomic counter case, state is passed as reference to
all the goroutines, and it is modified there atomically.

But, in some cases, such approach can be error prone such as
- having multiple channels or
- involving multiple mutexes.

In such cases, it's better to have just one goroutine with the state or a
stateful-goroutine,
and the other goroutines communicating read or writes through channels.
**/

package sync_examples

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// program state is a map[int]int

type readOp struct {
	key     int
	valResp chan int
}

type writeOp struct {
	key         int
	val         int
	successResp chan bool
}

func RunStatefulGoroutinesExample() {

	// we'll keep track of total read and write operations
	var (
		readOpsCount       uint64
		failedReadOpsCount uint64
		writeOpsCount      uint64
	)

	readsChan := make(chan readOp)
	writesChan := make(chan writeOp)
	doneChan := make(chan bool, 1)

	// stateful goroutine
	go func() {
		state := make(map[int]int)
	stateloop:
		for {
			select {
			case read := <-readsChan:
				{
					// fmt.Println("read request recd for", read.key)
					if val, exists := state[read.key]; exists {
						read.valResp <- val
					} else {
						read.valResp <- -1
					}

				}
			case write := <-writesChan:
				{
					// fmt.Println("write request recd", write.key, " -> ", write.val)
					state[write.key] = write.val
					write.successResp <- true
				}
			case <-doneChan:
				{
					fmt.Printf("state: %v\n", state)
					break stateloop
				}
			}
		}
	}()

	// start writer goroutines
	numWriters := 10
	for i := 0; i < numWriters; i++ {
		go writer(&writesChan, &writeOpsCount)
	}

	// start reader goroutines
	numReaders := 100
	for r := 0; r < numReaders; r++ {
		go reader(&readsChan, &readOpsCount, &failedReadOpsCount)
	}

	// let it run for some time
	t := time.NewTimer(time.Second * 1)
	<-t.C

	doneChan <- true
	t = time.NewTimer(time.Millisecond * 1)
	<-t.C

	readOpsFinal := atomic.LoadUint64(&readOpsCount)
	failedOpsFinal := atomic.LoadUint64(&failedReadOpsCount)
	writeOpsFinal := atomic.LoadUint64(&writeOpsCount)

	fmt.Printf("readOpsFinal: %v, writeOpsFinal: %v\n", readOpsFinal, writeOpsFinal)
	fmt.Printf("failedOpsFinal: %v\n", failedOpsFinal)
}

func writer(writeOpsChan *chan writeOp, writeOpsCount *uint64) {
	for {
		// create a write operation, with random key and val
		write := writeOp{
			key:         rand.Intn(5),
			val:         rand.Intn(100),
			successResp: make(chan bool),
		}

		// send write op
		*writeOpsChan <- write

		// block till success response is recd
		<-write.successResp
		// fmt.Println("----> write success")

		// count one successful write operation
		atomic.AddUint64(writeOpsCount, 1)

		// take some delay
		time.Sleep(time.Millisecond)
	}
}

func reader(readOpsChan *chan readOp, readOpsCount *uint64, failedReadOpsCount *uint64) {
	for {
		// create a read operation
		read := readOp{
			key:     rand.Intn(5),
			valResp: make(chan int),
		}

		// send read op
		*readOpsChan <- read

		// block till value is recd
		val := <-read.valResp
		if val < 0 {
			fmt.Println("failed key", read.key)
			atomic.AddUint64(failedReadOpsCount, 1)
		} else {
			// count one successful read operation
			atomic.AddUint64(readOpsCount, 1)
		}
		// fmt.Printf("<---- recd val: %v\n", val)

		// take some delay
		time.Sleep(time.Millisecond)

	}
}

package sync_examples

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const N_WORKERS = 1000
const ITERATIONS = 1000

// task:
// run `N_WORKERS` goroutines that each increments a counter `ITERATIONS` times

func RunAtomicCounter() {
	start := time.Now()
	withoutAtomic := run(false)
	nonAtomicTook := time.Since(start)
	fmt.Printf("withoutAtomic: %v\n", withoutAtomic)
	// fmt.Printf("took %v\n", nonAtomicTook)

	start = time.Now()
	withAtomic := run(true)
	atomicTook := time.Since(start)
	fmt.Printf("withAtomic: %v\n", withAtomic)
	// fmt.Printf("took %v\n", atomicTook)

	fmt.Printf("(atomicTook / nonAtomicTook): %d\n", (atomicTook / nonAtomicTook))
}

func run(isAtomic bool) uint64 {
	// unsigned integer (always positive)
	var counter uint64
	var wg sync.WaitGroup
	var incrementer func(*uint64)

	if isAtomic {
		incrementer = atomicIncrementer
	} else {
		incrementer = nonAtomicIncrementer
	}

	for i := 0; i < N_WORKERS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := 0; i < ITERATIONS; i++ {
				incrementer(&counter)
			}
		}()
	}
	wg.Wait()

	return counter
}

func atomicIncrementer(counter *uint64) {
	atomic.AddUint64(counter, 1)
}

func nonAtomicIncrementer(counter *uint64) {
	*counter++
}

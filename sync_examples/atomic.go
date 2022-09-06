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
	withoutAtomic := runWithoutAtomic()
	nonAtomicTook := time.Since(start)
	fmt.Printf("withoutAtomic: %v\n", withoutAtomic)
	// fmt.Printf("took %v\n", nonAtomicTook)

	start = time.Now()
	withAtomic := runWithAtomic()
	atomicTook := time.Since(start)
	fmt.Printf("withAtomic: %v\n", withAtomic)
	// fmt.Printf("took %v\n", atomicTook)

	fmt.Printf("(atomicTook / nonAtomicTook): %d\n", (atomicTook / nonAtomicTook))
}

func runWithoutAtomic() uint64 {
	// unsigned integer (always positive)
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < N_WORKERS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := 0; i < ITERATIONS; i++ {
				counter++
			}
		}()
	}
	wg.Wait()

	return counter
}

func runWithAtomic() uint64 {
	// unsigned integer (always positive)
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < N_WORKERS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := 0; i < ITERATIONS; i++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()

	return counter
}

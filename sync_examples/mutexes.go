package sync_examples

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[key]++
}

func RunMutexExample() {
	container := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	var wg sync.WaitGroup

	doIncrement := func(key string, iterations int) {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			container.increment(key)
		}
	}

	wg.Add(1)
	go doIncrement("a", 10000)

	wg.Add(1)
	go doIncrement("b", 10000)

	wg.Add(1)
	go doIncrement("c", 10000)

	wg.Add(1)
	go doIncrement("c", 10000)

	wg.Wait()
	fmt.Printf("container.counters: %v\n", container.counters)

}

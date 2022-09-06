// You can edit this code!
// Click here and start typing.
package timemachine

import (
	"fmt"
	"time"
)

type Limiter struct {
	maxCount int
	count    int
	ticker   *time.Ticker
	ch       chan struct{}
}

func (l *Limiter) run() {
	for {
		// if counter has reached 0: block until next tick
		if l.count <= 0 {
			<-l.ticker.C
			l.count = l.maxCount
		}

		// otherwise:
		select {
		case l.ch <- struct{}{}:
			l.count--

		case <-l.ticker.C:
			l.count = l.maxCount
		}
	}
}

func (l *Limiter) Wait() {
	d := <-l.ch
	fmt.Printf("d: %v\n", d)
}

func NewLimiter(d time.Duration, count int) *Limiter {
	l := &Limiter{
		maxCount: count,
		count:    count,
		ticker:   time.NewTicker(d),
		ch:       make(chan struct{}),
	}
	go l.run()

	return l
}

func bursts(l *Limiter) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		l.Wait()
		fmt.Println("burst ", i)
	}

	fmt.Println("--- bursts:", time.Since(start), "elapsed")
}

func slowIterations(l *Limiter) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		// <-time.After(250 * time.Millisecond)
		l.Wait()
		fmt.Println("iteration ", i)
	}

	fmt.Println("--- slowIterations:", time.Since(start), "elapsed")
}

func RunRateLimiting2() {
	fmt.Println("Hello, 世界")

	l := NewLimiter(time.Second, 5)

	bursts(l)
	slowIterations(l)
}

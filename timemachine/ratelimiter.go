package timemachine

import (
	"errors"
	"fmt"
	"time"
)

type RateLimiter struct {
	interval   time.Duration
	burstCount int
	count      int
	ticker     *time.Ticker
	C          chan time.Time
}

func (rl *RateLimiter) resetBurst() {
	fmt.Println("---- reset burst")
	rl.count = rl.burstCount
}

func (rl *RateLimiter) run() {
	// if burst count is neg or 0, it is considered as normal rate limiter (pass 1 event every d duration)
	if rl.burstCount <= 0 {
		for range rl.ticker.C {
			rl.C <- time.Now()
		}
	} else {
		// this is a burst limiter
		// i.e.
		// every d duration, allows `burstCount` events to pass and then wait for d duration

		// handle if count has become negative or 0

		for {
			if rl.count <= 0 {
				// wait till next timer
				fmt.Println("---- waiting for timer")
				<-rl.ticker.C
				rl.resetBurst()
			}

			rl.C <- time.Now()
			rl.count--
		}

	}
}

func (rl *RateLimiter) Wait() {
	<-rl.C
}

func NewRateLimiter(interval time.Duration) (*RateLimiter, error) {
	rl := &RateLimiter{
		interval: interval,
		ticker:   time.NewTicker(interval),
		C:        make(chan time.Time),
	}
	go rl.run()
	return rl, nil
}

func NewBurstyRateLimiter(interval time.Duration, burstCount int) (*RateLimiter, error) {
	if burstCount <= 0 {
		return nil, errors.New("burstCount should be > 0")
	}

	rl := &RateLimiter{
		interval:   interval,
		ticker:     time.NewTicker(interval),
		burstCount: burstCount,
		C:          make(chan time.Time, burstCount),
	}
	for i := 0; i < burstCount; i++ {
		rl.C <- time.Now()
	}
	go rl.run()
	return rl, nil
}

package timemachine

import (
	"errors"
	"time"
)

type RateLimiter struct {
	interval time.Duration
	isBurst  bool
	ticker   *time.Ticker
	C        chan time.Time
}

func (rl *RateLimiter) run() {
	for {
		rl.C <- <-rl.ticker.C
	}
}

func (rl *RateLimiter) Wait() time.Time {
	return <-rl.C
}

func NewRateLimiter(interval time.Duration) (*RateLimiter, error) {
	rl := &RateLimiter{
		interval: interval,
		isBurst:  false,
		ticker:   time.NewTicker(interval),
		C:        make(chan time.Time),
	}
	go rl.run()
	return rl, nil
}

func NewBurstyRateLimiter(interval time.Duration, isBurst bool, burstCount int) (*RateLimiter, error) {
	if burstCount <= 0 {
		return nil, errors.New("burstCount should be > 0")
	}
	rl := &RateLimiter{interval: interval, isBurst: isBurst}
	rl.C = make(chan time.Time, burstCount)
	for i := 0; i < burstCount; i++ {
		rl.C <- time.Now()
	}
	return rl, nil
}

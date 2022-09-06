package timemachine

import (
	"fmt"
	"time"
)

type Request struct {
	data int
}

const N_REQUESTS = 5
const INTERVAL = time.Millisecond * 1000

func sendAndReceiveRequests(n_requests int, rateLimiter *RateLimiter) {
	// create and send requests
	requestsChannel := make(chan Request, n_requests)
	for i := 0; i < n_requests; i++ {
		requestsChannel <- Request{data: i}
	}
	close(requestsChannel)

	// receive requests
	var start time.Time
	for request := range requestsChannel {
		start = time.Now()
		// wait for rate limit channel
		rateLimiter.Wait()
		fmt.Printf("request recd: %v, took: %v\n", request, time.Since(start))
	}
}

func runNormalLimiter() {
	fmt.Println("---- Running Normal Rate Limiter with interval: ", INTERVAL)

	// create rate limiter instance
	if rateLimiter, err := NewRateLimiter(INTERVAL); err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		sendAndReceiveRequests(N_REQUESTS, rateLimiter)
	}

}

func runBurstyLimiter() {
	fmt.Println("---- Running Bursty Rate Limiter with interval: ", INTERVAL)

	n_requests := N_REQUESTS * 3
	// create rate limiter instance
	if rateLimiter, err := NewBurstyRateLimiter(INTERVAL, 3); err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		sendAndReceiveRequests(n_requests, rateLimiter)
	}

}

func RunRateLimiting() {
	runNormalLimiter()
	runBurstyLimiter()
}

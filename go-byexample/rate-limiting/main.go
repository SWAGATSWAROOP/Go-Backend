package main

import (
	"fmt"
	"time"
)

// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.

func main() {
	requests := make(chan int, 5)
	// First we’ll look at basic rate limiting. Suppose we want to limit our handling of incoming requests. We’ll serve these requests off a channel of the same name.
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())

		// We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit. We can accomplish this by buffering our limiter channel. This burstyLimiter channel will allow bursts of up to 3 events.
		burstyLimiter := make(chan time.Time, 3)

		// Fill up the channel to represent allowed bursting.
		for range 3 {
			burstyLimiter <- time.Now()
		}

		// Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
		go func() {
			for t := range time.Tick(200 * time.Millisecond) {
				burstyLimiter <- t
			}
		}()
		// IMP Every 200ms, it adds one more token to the channel — only if it's not already full (capacity = 3).

		// Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of burstyLimiter.
		burstyRequests := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			burstyRequests <- i
		}
		close(burstyRequests)
		// For the second batch of requests we serve the first 3 immediately because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.
		for req := range burstyRequests {
			<-burstyLimiter
			fmt.Println("request", req, time.Now())
		}
		/*First 3 requests: pass immediately (tokens are already in the channel).
		Next 2: have to wait ~200ms each, as tokens are added to the channel every 200ms by the goroutine.
		*/
	}
}

/*
Bursty
🧠 Visual Summary:
Request	Time it executes	Why
1	immediately	token already present
2	immediately	token already present
3	immediately	token already present
4	after ~200ms	waited for next token
5	after ~400ms	waited for another token
*/

package main

import (
	"fmt"
	"time"
)

// Ticker : tickers are for when you want to do something repeatedly at regular intervals.
func Ticker() {
	ticker := time.NewTicker(500 * time.Millisecond)

	//	use the range builtin on the channel to iterate over the values as they arrive every 500ms.
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at (microsecond)", t)
		}
	}()

	//Once a ticker is stopped it won’t receive any more values on its channel.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop() // Stops after 1600ms.
	fmt.Println("Ticker stopped")
}

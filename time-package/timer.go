package main

import (
	"fmt"
	"time"
)

// Timer : Timers are for when we want to do something once in the future
func Timer() {
	// Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second) // tell timer to wait 2 seconds.

	// blocks on the timer’s channel C until it sends a value indicating that the timer expired.
	<-timer1.C
	fmt.Println("Timer 1 expired after 2 seconds") //  first timer will expire ~2s after we start the program

	// One reason a timer may be useful is that you can cancel the timer before it expires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

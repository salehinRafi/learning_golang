package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to channel c
}

/*

GoChannel :

a. What is Channel?
Channels allow us to pass data (send & receive) between go routines with the channel operator, <-

b. We can create channel based on:-
	i. Create channel (without buffer)
		_element_ := make(chan _type_)

		- By default, sends and receives block wait until the other side is ready.
		- This allows goroutines to synchronize without explicit locks or condition variables.

	ii. Create channel with buffer
		_element_ := make(chan _type_, _buffer-length_)

		- Sends to a buffered channel block only when the buffer is full.
		- Receives block when the buffer is empty.

Extra Note:
i. Send & Receive data (using operator <-)

	ch <- v
		- Send v to channel ch.

	v := <-ch
		- Receive from ch, and assign value to v.

ii. Beware of deadlocks!!
	How it happens:-
	-  end the buffer without letting the code a chance to read/remove a value from the channel.

*/
func GoChannel() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) // declaration and initialization
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c and assign to x, y
	fmt.Println(x, y, x+y)
}

/*
Range and Close in Channel Goroutine

	i. Sender
		-	A sender can close a channel to indicate that no more values will be sent.
	ii. Receiver
		-	Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression.
	iii. Expression
		- v, ok := <-ch , where v as a sender, ok as a receiver. (ok is false if there are no more values to receive and the channel is closed.)

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren’t like files; you don’t usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // Closing a channel indicates that no more values will be sent on it.
}

func GoRangeClose() {

	// 1st Example
	// like maps and slices, channels must be created before use:
	c := make(chan int, 10)
	go fibonacci(cap(c), c) //cap built-in function returns the capacity of v, according to its type. In this case return as channel type
	for i := range c {      //receives values from the channel repeatedly until it is closed.
		fmt.Println(i)
	}

	// 2nd Example
	queue := make(chan string, 2)
	//iterate over 2 values in the queue channel.
	queue <- "one"
	queue <- "two"
	close(queue) // Closing a channel indicates that no more values will be sent on it.

	// range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}

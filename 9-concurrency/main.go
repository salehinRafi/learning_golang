package main

import "fmt"

func main() {
	/*
		Go encourages a different approach in which shared values are passed around on channels and, in fact, never actively shared by separate threads of execution.

		Only one goroutine has access to the value at any given time.

		Go approach:-
		-	Do not communicate by sharing memory; instead, share memory by communicating.

	*/
	fmt.Println("------Goroutines Sections------")
	Goroutines()

	fmt.Println("------Channel Sections------")
	GoChannel()

	fmt.Println("------Range & Close Sections------")
	GoRangeClose()

	fmt.Println("------Select Sections------")
	GoSelect()
}

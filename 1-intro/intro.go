// A comment

/*
Multiline Comment
*/

// Every Go program starts with a package declaration which provides a way for
// use to reuse code
package main

// import allows use to use code from other packages
import (
	"fmt" // The format package provides formatting for input and output
)

// Functions start with func and surround the code with { }
// main is the function that is executed when you execute your program
func main() {
	//The Println function of the fmt package is used to write text to the standard string output, which
	// is surrounded by double quotes and a newline to the screen
	fmt.Println("Hello World")

	// You can get a description of a function by typing `godoc fmt Println`
	// for example in the terminal
}

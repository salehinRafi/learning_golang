package main

import (
	"fmt"
)

// Variable : Variable declaration in Go
func Variable() {

	// Variables are statically typed, which means their type can't change (immutable
	// Variable names must start with a letter and may contain letters, numbers or the _

	// declaration without initialization
	var a int
	a = 1
	fmt.Println(a)

	// declaration with initialization
	var b int = 42
	fmt.Println(b)

	// type omitted, will be inferred
	var e = 42
	fmt.Println(e)

	// shorthand, only in func bodies, omit var keyword, type is always implicit
	foo := 42
	fmt.Println(foo)

	// A constant is a variable with a value that can't be changed. Constants can be character, string, boolean, or numeric values.
	const constant = "This is a constant"

	// Multiple declaration at once
	var c, d int = 42, 1302
	fmt.Println(c, d)

	var (
		varA = 2
		varB = "Hellow"
	)
	fmt.Println(varA, varB)

}

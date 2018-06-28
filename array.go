package main

import "fmt"

func Array() {

	fmt.Println("------Array Sections------")
	var a [2]string

	a[0] = "Salehin"
	a[1] = "Rafi"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var b [2]int
	b[0] = 1
	b[1] = 5

	fmt.Println(b)
}

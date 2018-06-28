package main

import "fmt"

func Function() {
	fmt.Println("------Function Sections------")
	var a = 10
	var b = 20
	var c int
	c = add(a, b)

	fmt.Println("Addition function result:", c)

}

func add(num1, num2 int) int {
	var result int
	result = num1 + num2
	return result
}

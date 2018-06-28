package main

import (
	"fmt"
	"syscall"
)

func EnvironmentVariable() {
	fmt.Println("------Environment Variable Sections------")
	env := syscall.Environ()

	for i := range env {
		fmt.Println(env[i])
	}
}

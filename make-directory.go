package main

import (
	"fmt"
	"syscall"
)

func MakeDirectory() {
	fmt.Println("------Make Directory Sections------")
	err := syscall.Mkdir("new-dir", 0754)

	if err == nil {
		fmt.Println("Directory Created")
	}
}

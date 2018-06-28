package main

import (
	"fmt"
	"syscall"
)

func GetProcessID() {
	fmt.Println("------Get Process ID Sections------")
	pid := syscall.Getpid()
	fmt.Println(pid)
}

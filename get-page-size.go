package main

import (
	"fmt"
	"os"
)

func GetPageSize() {
	fmt.Println("------Get Page Size Sections------")
	pagesize := os.Getpagesize()
	fmt.Println(pagesize)
}

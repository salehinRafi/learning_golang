package main

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02T"
)

func DiffDates() {
	v := "2018-09-08"
	then, err := time.Parse(timeFormat, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	duration := time.Since(then)
	fmt.Println(duration.Hours())
	fmt.Println(duration.Minutes())

}

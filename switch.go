package main

import (
	"fmt"
)

func Switch() {
	fmt.Println("------Switch Sections------")
	var i = 2
	switch i {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	}

}

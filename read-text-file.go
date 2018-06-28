package main

import (
	"fmt"
	"io/ioutil"
)

func ReadTextFile() {
	fmt.Println("------Read Text File Sections------")
	b, err := ioutil.ReadFile("ReadTextFile.txt")

	if err != nil {
		fmt.Println(err)
	}

	str := string(b)
	fmt.Println(str)
}

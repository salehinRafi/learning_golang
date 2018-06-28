package main

import (
	"fmt"
	"log"
)

func GenerateErrorLog() {
	fmt.Println("------Generate Error Log Sections------")
	log.Fatalln("Fatal Error Example") // Any Fatal Error occur system will exit
}

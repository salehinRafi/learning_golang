package main

import (
	"fmt"
	"log"
	"os"
)

func WriteLogToFile() {
	fmt.Println("------Write Log To File Sections------")
	file, err := os.OpenFile("WriteLogToFile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed")
	}
	log.SetOutput(file)
	log.Panicln("Log To File")
}

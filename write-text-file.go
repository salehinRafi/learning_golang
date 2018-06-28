package main

import (
	"fmt"
	"log"
	"os"
)

func WriteTextFile() {
	fmt.Println("------Write Text File Sections------")
	file, err := os.Create("WriteTextFile.txt")
	if err != nil {
		log.Fatal("We got error", err)
	}

	defer file.Close()
	fmt.Fprint(file, "Salehin Rafi Go Lang Tutorial")
}

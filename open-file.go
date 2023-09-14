package main

import (
	"fmt"
	"os"
)

func OpenFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print(fmt.Errorf("error opening file: %v", err))
		os.Exit(1)
	}

	return file
}

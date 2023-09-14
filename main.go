package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	commandsToExecute := map[string]interface{}{
		"-c": readBytes,
		"-l": countLines,
		"-w": countWords,
		"-m": countCharacteres,
	}

	/* os.Args provides access to raw command-line arguments.
	Note that the first value in this slice is the path to the program,
	and os.Args[1:] holds the arguments to the program. */
	args := os.Args[1:]

	if len(args) == 1 {
		fileName := args[0]
		defaultNoCommandProvide(fileName)
	} else {
		command := args[0]
		fileName := args[1]

		//get the respective function to execute
		if f, ok := commandsToExecute[command]; ok {
			value := f.(func(string) int)(fileName) //casting f to function and execute
			fmt.Printf("%8d %s\n", value, fileName)
		} else {
			fmt.Print(fmt.Errorf("invalid command param %v", command))
		}
	}
}

func defaultNoCommandProvide(fileName string) {
	totalBytes := readBytes(fileName)
	totalLines := countLines(fileName)
	totalWords := countWords(fileName)

	fmt.Printf("%8d %8d %8d %s\n", totalBytes, totalLines, totalWords, fileName)
}

func readBytes(fileName string) int {
	//Read file
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(fmt.Errorf("error reading file: %v", err))
		os.Exit(1)
	}

	//Count the number of bytes
	byteCount := len(fileContent)

	return byteCount
}

func countLines(fileName string) int {
	file := OpenFile(fileName)
	defer file.Close()

	linesCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(fmt.Errorf("error reading a file: %v", err))
		os.Exit(1)
	}

	return linesCount
}

func countWords(fileName string) int {
	file := OpenFile(fileName)
	defer file.Close()

	wordsCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordsCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(fmt.Errorf("error reading a file: %v", err))
		os.Exit(1)
	}

	return wordsCount
}

func countCharacteres(fileName string) int {
	file := OpenFile(fileName)
	defer file.Close()
	characteresCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			characteresCount += utf8.RuneCountInString(word)
		}
	}

	return characteresCount
}

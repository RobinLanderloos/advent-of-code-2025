package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func EnumerateFile(path string) <-chan string {
	ch := make(chan string)

	go enumerateFile(path, ch)

	return ch
}

func enumerateFile(path string, ch chan string) {
	defer close(ch)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		ch <- line
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}

func ReadLines(path string) []string {
	file, err := os.Open(path)
	lines := make([]string, 0)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return lines
}

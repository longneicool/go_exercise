package main

import (
	"bufio"
	"fmt"
	"os"
)

func readWord(path string) map[string]int {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file %s\n", path)
		os.Exit(1)
	}

	letters := make(map[string]int)
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		letters[scanner.Text()]++
	}

	return letters
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s file", os.Args[0])
	}

	letters := readWord(os.Args[1])
	fmt.Println(letters)
}

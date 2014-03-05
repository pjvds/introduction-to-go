package main

import (
	"bufio"
	"os"
)

func CountLines(filename string) (lines int) {
	file := os.Open(filename)
	defer file.Close() // HL3

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}

	return
}

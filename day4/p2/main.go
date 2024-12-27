package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var TARGET_WORD string = "MAS"
var REVERSED_TARGET_WORD string = "SAM"

func main() {
	// Open File
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("error reading file")
	}
	defer file.Close()

	// Process File
	matrix := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		newRow := []string{}
		for _, r := range line {
			newRow = append(newRow, string(r))
		}
		matrix = append(matrix, newRow)
	}

	total := 0

	// Compute Result
	for r := 0; r <= len(matrix)-len(TARGET_WORD); r++ {
		for c := 0; c <= len(matrix[r])-len(TARGET_WORD); c++ {
			if matrix[r+1][c+1] != "A" {
				continue
			}
			leftString := ""
			for dist := 0; dist < len(TARGET_WORD); dist++ {
				leftString += matrix[r+dist][c+dist]
			}
			rightString := ""
			for dist := 0; dist < len(TARGET_WORD); dist++ {
				rightString += matrix[r+dist][c+len(TARGET_WORD)-1-dist]
			}
			leftStringMatches := leftString == TARGET_WORD || leftString == REVERSED_TARGET_WORD
			rightStringMatches := rightString == TARGET_WORD || rightString == REVERSED_TARGET_WORD
			if leftStringMatches && rightStringMatches {
				total += 1
			}
		}
	}

	fmt.Println(total)
}

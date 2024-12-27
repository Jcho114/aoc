package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var TARGET_WORD string = "XMAS"
var DIRECTIONS [][]int = [][]int{
	{1, 0}, {0, 1},
	{1, 1}, {-1, 0},
	{0, -1}, {-1, -1},
	{1, -1}, {-1, 1},
}

func main() {
	// Read File
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("unable to read file")
	}
	defer file.Close()

	// Process Input
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

	// Word Search
	total := 0
	for r, line := range matrix {
		for c := range line {
			if matrix[r][c] == string(TARGET_WORD[0]) {
				for _, direction := range DIRECTIONS {
					s := string(matrix[r][c])
					di, dj := direction[0], direction[1]
					for ci, cj, count := r+di, c+dj, 1; ci >= 0 && cj >= 0 && ci < len(matrix) && cj < len(matrix[ci]) && count < len(TARGET_WORD); ci, cj, count = ci+di, cj+dj, count+1 {
						s += matrix[ci][cj]
					}
					if s == TARGET_WORD {
						total += 1
					}
				}
			}
		}
	}
	fmt.Println(total)
}

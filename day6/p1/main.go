package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Key struct {
	Row int
	Col int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("error opening input file")
	}
	defer file.Close()

	matrix := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	// Find guard position
	var gr int
	var gc int
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == '^' {
				gr, gc = r, c
				break
			}
		}
	}

	N, M := len(matrix), len(matrix[0])
	directions := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	currDirection := 0
	res := make(map[Key]bool)
	for {
		key := Key{Row: gr, Col: gc}
		if _, ok := res[key]; !ok {
			res[key] = true
		}
		direction := directions[currDirection]
		dr, dc := direction[0], direction[1]
		nr, nc := gr+dr, gc+dc
		if nr == -1 || nc == -1 || nr == N || nc == M {
			break
		} else if matrix[nr][nc] == '#' {
			currDirection = (currDirection + 1) % len(directions)
		} else {
			gr, gc = nr, nc
		}
	}

	fmt.Println(len(res))
}

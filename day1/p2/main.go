package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftNumbers []int
	rightNumberCounts := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		leftNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalf("error parsing line")
		}
		rightNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("error parsing line")
		}
		leftNumbers = append(leftNumbers, leftNumber)
		_, exists := rightNumberCounts[rightNumber]
		if !exists {
			rightNumberCounts[rightNumber] = 0
		}
		rightNumberCounts[rightNumber] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file")
	}

	total := 0
	for _, leftNumber := range leftNumbers {
		_, exists := rightNumberCounts[leftNumber]
		if exists {
			total += leftNumber * rightNumberCounts[leftNumber]
		}
	}

	fmt.Println(total)
}

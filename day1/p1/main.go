package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var rightNumbers []int

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
		rightNumbers = append(rightNumbers, rightNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file")
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	totalDifference := 0
	for i := 0; i < len(leftNumbers); i++ {
		leftNumber := leftNumbers[i]
		rightNumber := rightNumbers[i]
		var difference int
		if leftNumber < rightNumber {
			difference = rightNumber - leftNumber
		} else {
			difference = leftNumber - rightNumber
		}
		totalDifference += difference
	}

	fmt.Println(totalDifference)
}

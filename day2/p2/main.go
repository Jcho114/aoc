package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(report []int) bool {
	checkDecreasing := report[1] < report[0]
	for i := 1; i < len(report); i++ {
		if checkDecreasing && report[i] >= report[i-1] {
			return false
		}
		if !checkDecreasing && report[i] <= report[i-1] {
			return false
		}
		var difference int
		if report[i] > report[i-1] {
			difference = report[i] - report[i-1]
		} else {
			difference = report[i-1] - report[i]
		}
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

func isReportSomewhatSafe(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		slicedNumbers := make([]int, 0, len(numbers)-1)
		slicedNumbers = append(slicedNumbers, numbers[:i]...)
		slicedNumbers = append(slicedNumbers, numbers[i+1:]...)
		if isReportSafe(slicedNumbers) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("error reading file")
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringNumbers := strings.Split(line, " ")
		var numbers []int
		for i := 0; i < len(stringNumbers); i++ {
			number, err := strconv.Atoi(stringNumbers[i])
			if err != nil {
				log.Fatalf("error in parsing line")
			}
			numbers = append(numbers, number)
		}
		if isReportSomewhatSafe(numbers) {
			total += 1
		}
	}

	fmt.Println(total)
}

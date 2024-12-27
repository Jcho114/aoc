package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("error reading file")
	}
	defer file.Close()

	generalRegex, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|don't\(\)|do\(\)`)
	multiplyRegex, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	total := 0
	linesCombined := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		linesCombined += line
	}
	matches := generalRegex.FindAllString(linesCombined, -1)

	shouldProcess := true
	for _, match := range matches {
		if match == "don't()" {
			shouldProcess = false
		} else if match == "do()" {
			shouldProcess = true
		} else if shouldProcess {
			captures := multiplyRegex.FindStringSubmatch(match)
			n1, err := strconv.Atoi(captures[1])
			if err != nil {
				log.Fatalf("error in parsing number")
			}
			n2, err := strconv.Atoi(captures[2])
			if err != nil {
				log.Fatalf("error in parsing number")
			}
			product := n1 * n2
			total += product
		}
	}

	fmt.Println(total)
}

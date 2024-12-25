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

	r, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	total := 0
	linesCombined := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		linesCombined += line
	}
	matches := r.FindAllString(linesCombined, -1)
	for _, match := range matches {
		captures := r.FindStringSubmatch(match)
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

	fmt.Println(total)
}

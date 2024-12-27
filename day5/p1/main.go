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
	// Open File
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("unable to open file")
	}
	defer file.Close()

	// Process Page Rules
	rules := make(map[int]map[int]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		pages := strings.Split(line, "|")
		src, err := strconv.Atoi(pages[0])
		if err != nil {
			log.Fatalf("error in parsing integer")
		}
		dest, err := strconv.Atoi(pages[1])
		if err != nil {
			log.Fatalf("error in parsing integer")
		}
		if _, ok := rules[src]; !ok {
			rules[src] = make(map[int]bool)
		}
		rules[src][dest] = true
	}

	// Process Updates
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		pages := strings.Split(line, ",")
		processed := []int{}
		isValidUpdate := true
		for _, rawPage := range pages {
			page, err := strconv.Atoi(rawPage)
			if err != nil {
				log.Fatalf("error in parsing integer")
			}
			for _, processedPage := range processed {
				if _, ok := rules[page][processedPage]; ok {
					isValidUpdate = false
				}
			}
			if !isValidUpdate {
				break
			}
			processed = append(processed, page)
		}
		if isValidUpdate {
			centerIndex := len(processed) / 2
			total += processed[centerIndex]
		}
	}

	fmt.Println(total)
}

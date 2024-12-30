package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func toposort(graph map[int]map[int]bool, nodes []int) []int {
	result := []int{}

	// Create set for O(1) node check
	nodeSet := make(map[int]bool)
	for _, node := range nodes {
		nodeSet[node] = true
	}

	// Indegree Preprocessing
	indegrees := make(map[int]int)
	for _, node := range nodes {
		indegrees[node] = 0
	}
	for src, dests := range graph {
		for dest := range dests {
			if _, ok := nodeSet[src]; ok {
				if _, ok := nodeSet[dest]; ok {
					indegrees[dest] += 1
				}
			}
		}
	}

	// Init Queue
	queue := []int{}
	for node, indegree := range indegrees {
		if indegree == 0 {
			queue = append(queue, node)
		}
	}

	// Kahn Algorithm
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for dest := range graph[node] {
			indegrees[dest] -= 1
			if indegrees[dest] == 0 {
				queue = append(queue, dest)
			}
		}

		result = append(result, node)
	}

	return result
}

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
			processed = append(processed, page)
		}
		if !isValidUpdate {
			processed = toposort(rules, processed)
			centerIndex := len(processed) / 2
			total += processed[centerIndex]
		}
	}

	fmt.Println(total)
}

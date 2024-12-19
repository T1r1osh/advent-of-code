package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("solutions/2024/Day19/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	index := 0
	var available_patterns []string
	designs := []string{}
	for scanner.Scan() {

		line := scanner.Text()
		if index == 0 {
			available_patterns = strings.Split(line, ", ")
		}
		if index > 1 {
			designs = append(designs, line)
		}
		index++
	}
	//log.Println(available_patterns)
	//log.Println(designs)
	log.Println(Task1(designs, available_patterns))
	log.Println(Task2(designs, available_patterns))
}

func Task1(designs, available_patterns []string) int {
	sum := 0
	for _, design := range designs {
		dp := make([]bool, len(design)+1)
		dp[0] = true
		for j := 0; j < len(design); j++ {
			if dp[j] {
				for _, available := range available_patterns {
					substr := Substring(design, j, len(available))

					if (j+len(available)) <= len(design) &&
						available == substr {
						dp[j+len(available)] = true
					}
				}
			}
		}
		if dp[len(design)] {
			sum++
		}
	}
	return sum
}

func Task2(designs, available_patterns []string) int {
	sum := 0
	for _, design := range designs {
		sum += countWaysToConstruct(design, available_patterns)
	}
	return sum
}

func countWaysToConstruct(design string, availablePatterns []string) int {
	dp := make([]int, len(design)+1)
	dp[0] = 1

	for j := 0; j < len(design); j++ {
		if dp[j] > 0 {
			for _, pattern := range availablePatterns {
				if j+len(pattern) <= len(design) && design[j:j+len(pattern)] == pattern {
					dp[j+len(pattern)] += dp[j]
				}
			}
		}
	}

	return dp[len(design)]
}

func Substring(s string, start, length int) string {
	if start < 0 || start >= len(s) || length < 0 || start+length > len(s) {
		return "" // invalid range
	}
	return s[start : start+length]
}

package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func task1(in string) int {

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(in, -1)

	total := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		total += x * y
	}
	return total
}

func task2(in string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(in, -1)
	total := 0
	enable := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			enable = true
		case "don't()":
			enable = false
		default:
			if enable {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				total += x * y
			}
		}
	}
	return total
}

func main() {

	file, _ := os.Open("solutions/2024/Day03/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	input := sb.String()

	println("Part one answer is:", task1(input))
	println("Part two answer is:", task2(input))

}

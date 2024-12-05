package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules := make(map[int][]int, 0)

	ruleFile, _ := os.Open("solutions/2024/Day05/rules.txt")
	defer ruleFile.Close()
	scanner := bufio.NewScanner(ruleFile)

	for scanner.Scan() {
		line := strArrToIntArr(strings.Split(scanner.Text(), "|"))
		rules[line[0]] = append(rules[line[0]], line[1])
	}

	inputFile, _ := os.Open("solutions/2024/Day05/input.txt")
	defer inputFile.Close()
	scanner = bufio.NewScanner(inputFile)
	sum := 0
	sumnok := 0
	for scanner.Scan() {
		line := strArrToIntArr(strings.Split(scanner.Text(), ","))
		if correctInput(line, rules) {
			sum = sum + Task1(line)
		} else {
			sumnok = sumnok + Task2(line, rules)
		}

	}
	println("Task 1: ", sum)
	println("Task 2: ", sumnok)
}

func Task1(line []int) int {
	return line[len(line)/2]
}

func Task2(line []int, rules map[int][]int) int {
	correctLine := make([]int, len(line))
	for _, val := range line {
		counter := 0
		for _, forbid := range rules[val] {
			if slices.Index(line, forbid) != -1 {
				counter++
			}
		}
		correctLine[counter] = val
	}

	return correctLine[len(correctLine)/2]
}

func strArrToIntArr(arr []string) []int {
	intarr := make([]int, len(arr))
	for i, val := range arr {
		intarr[i], _ = strconv.Atoi(val)
	}
	return intarr
}

func correctInput(line []int, rules map[int][]int) bool {
	for index, val := range line {
		for i := 0; i < index; i++ {
			for j := 0; j < len(rules[val]); j++ {
				if line[i] == rules[val][j] {
					return false
				}
			}
		}
	}
	return true
}

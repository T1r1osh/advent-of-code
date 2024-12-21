package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var numPad map[string][2]int
var dirPad map[string][2]int

func init() {
	numPad = make(map[string][2]int, 12)
	numPad["0"] = [2]int{1, 0}
	numPad["A"] = [2]int{2, 0}
	numPad["1"] = [2]int{0, 1}
	numPad["2"] = [2]int{1, 1}
	numPad["3"] = [2]int{2, 1}
	numPad["4"] = [2]int{0, 2}
	numPad["5"] = [2]int{1, 2}
	numPad["6"] = [2]int{2, 2}
	numPad["7"] = [2]int{0, 3}
	numPad["8"] = [2]int{1, 3}
	numPad["9"] = [2]int{2, 3}

	dirPad = make(map[string][2]int, 6)
	dirPad["^"] = [2]int{1, 1}
	dirPad["A"] = [2]int{2, 1}
	dirPad["<"] = [2]int{0, 0}
	dirPad["v"] = [2]int{1, 0}
	dirPad[">"] = [2]int{2, 0}
}

func main() {
	file, _ := os.Open("solutions/2024/Day21/input.txt")
	defer file.Close()

	codes := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, line)
	}

	log.Println(Task1(codes))
	log.Println(Task2(codes))
}

func Task1(codes []string) int {
	sum := 0
	for _, code := range codes {
		sequence := dirPadToDirPad(dirPadToDirPad((numPadToKeyPad(code))))
		sum += convertCodeToNum(code) * len(sequence)
	}

	return sum
}

func Task2(codes []string) int {
	sum := 0
	cache := make(map[string][]int)
	for _, code := range codes {
		step1 := numPadToKeyPad(code)
		step2 := getCountAfterRobot(step1, 26, 1, cache)
		sum += step2 * convertCodeToNum(code)
	}

	return sum
}

//+---+---+---+
//| 7 | 8 | 9 |
//+---+---+---+
//| 4 | 5 | 6 |
//+---+---+---+
//| 1 | 2 | 3 |
//+---+---+---+
// 	  | 0 | A |
//+---+---+---+

func numPadToKeyPad(code string) string {
	seq := []string{}
	codeAsArray := strings.Split(code, "")
	start := numPad["A"]
	for _, char := range codeAsArray {
		goal := numPad[char]
		distance := [2]int{goal[0] - start[0], goal[1] - start[1]}

		moveX := []string{}
		for i := 0; i < abs(distance[0]); i++ {
			if distance[0] >= 0 {
				moveX = append(moveX, ">")
			} else {
				moveX = append(moveX, "<")
			}
		}

		moveY := []string{}
		for i := 0; i < abs(distance[1]); i++ {
			if distance[1] >= 0 {
				moveY = append(moveY, "^")
			} else {
				moveY = append(moveY, "v")
			}
		}

		// 1. moving with least turns
		// 2. moving < over ^ over v over >

		if start[1] == 0 && goal[0] == 0 {
			seq = append(seq, moveY...)
			seq = append(seq, moveX...)
		} else if start[0] == 0 && goal[1] == 0 {
			seq = append(seq, moveX...)
			seq = append(seq, moveY...)
		} else if distance[0] < 0 {
			seq = append(seq, moveX...)
			seq = append(seq, moveY...)
		} else if distance[0] >= 0 {
			seq = append(seq, moveY...)
			seq = append(seq, moveX...)
		}

		start = goal
		seq = append(seq, "A")

	}
	return strings.Join(seq[:], "")
}

//+---+---+---+
//|   | ^ | A |
//+---+---+---+
//| < | v | > |
//+---+---+---+

func dirPadToDirPad(code string) string {
	seq := []string{}
	codeAsArray := strings.Split(code, "")
	start := dirPad["A"]
	for _, char := range codeAsArray {
		goal := dirPad[char]
		distance := [2]int{goal[0] - start[0], goal[1] - start[1]}

		moveX := []string{}
		for i := 0; i < abs(distance[0]); i++ {
			if distance[0] >= 0 {
				moveX = append(moveX, ">")
			} else {
				moveX = append(moveX, "<")
			}
		}

		moveY := []string{}
		for i := 0; i < abs(distance[1]); i++ {
			if distance[1] >= 0 {
				moveY = append(moveY, "^")
			} else {
				moveY = append(moveY, "v")
			}
		}

		// 1. moving with least turns
		// 2. moving < over ^ over v over >

		if start[0] == 0 && goal[1] == 1 {
			seq = append(seq, moveX...)
			seq = append(seq, moveY...)
		} else if start[1] == 1 && goal[0] == 0 {
			seq = append(seq, moveY...)
			seq = append(seq, moveX...)
		} else if distance[0] < 0 {
			seq = append(seq, moveX...)
			seq = append(seq, moveY...)
		} else if distance[0] >= 0 {
			seq = append(seq, moveY...)
			seq = append(seq, moveX...)
		}

		start = goal
		seq = append(seq, "A")

	}
	return strings.Join(seq[:], "")
}

func getCountAfterRobot(seq string, maxRobots, robot int, cache map[string][]int) int {
	if val, ok := cache[seq]; ok {
		if val[robot-1] != 0 {
			return val[robot-1]
		}
	} else {
		cache[seq] = make([]int, maxRobots)
	}

	steps := dirPadToDirPad(seq)
	cache[seq][0] = len(steps)

	if robot == maxRobots {
		return len(seq)
	}

	splitSeq := getIndividualSteps(steps)
	count := 0
	for _, s := range splitSeq {
		// calculate small chunks
		c := getCountAfterRobot(s, maxRobots, robot+1, cache)
		if _, ok := cache[s]; !ok {
			cache[s] = make([]int, maxRobots)
		}
		cache[s][0] = c
		count += c
	}
	cache[seq][robot-1] = count
	return count

}

func getIndividualSteps(input string) []string {
	inputArray := strings.Split(input, "")
	output := []string{}
	current := []string{}

	for _, char := range inputArray {
		current = append(current, char)
		if char == "A" {
			output = append(output, strings.Join(current, ""))
			current = []string{}
		}
	}
	return output
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func convertCodeToNum(code string) int {
	codeNum, _ := strconv.Atoi(code[:3])
	return codeNum
}

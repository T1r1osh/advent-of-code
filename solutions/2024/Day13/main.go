package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Machine struct {
	A     [2]int
	B     [2]int
	Prize [2]int
}

func main() {
	file, err := os.Open("solutions/2024/Day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buttonRegex := regexp.MustCompile(`Button (A|B): X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	//var graph [][]int
	var machines []Machine
	var current Machine
	for scanner.Scan() {
		line := scanner.Text()

		if buttonMatch := buttonRegex.FindStringSubmatch(line); buttonMatch != nil {
			x, _ := strconv.Atoi(buttonMatch[2])
			y, _ := strconv.Atoi(buttonMatch[3])
			if buttonMatch[1] == "A" {
				current.A = [2]int{x, y}
			} else {
				current.B = [2]int{x, y}
			}
		}

		// Parse Prize line
		if prizeMatch := prizeRegex.FindStringSubmatch(line); prizeMatch != nil {
			x, _ := strconv.Atoi(prizeMatch[1])
			y, _ := strconv.Atoi(prizeMatch[2])
			current.Prize = [2]int{x, y}
			machines = append(machines, current)
			current = Machine{} // Reset for next block
		}
	}

	log.Println(Task1(machines))
	log.Println(Task2(machines))

}

func Task1(machines []Machine) int {

	tokens := 0
	for _, machine := range machines {
		// Prize = tokenA * A + tokenB * B
		tokenA := (machine.Prize[0]*machine.B[1] - machine.Prize[1]*machine.B[0]) / (machine.A[0]*machine.B[1] - machine.A[1]*machine.B[0])
		tokenB := (machine.A[0]*machine.Prize[1] - machine.A[1]*machine.Prize[0]) / (machine.A[0]*machine.B[1] - machine.A[1]*machine.B[0])

		if tokenA*machine.A[0]+tokenB*machine.B[0] == machine.Prize[0] && tokenA*machine.A[1]+tokenB*machine.B[1] == machine.Prize[1] {
			tokens += (tokenA * 3) + tokenB
		}
	}
	return tokens
}

func Task2(machines []Machine) int {

	tokens := 0
	for _, machine := range machines {
		machine.Prize[0], machine.Prize[1] = machine.Prize[0]+10000000000000, machine.Prize[1]+10000000000000
		// Prize = tokenA * A + tokenB * B
		tokenA := (machine.Prize[0]*machine.B[1] - machine.Prize[1]*machine.B[0]) / (machine.A[0]*machine.B[1] - machine.A[1]*machine.B[0])
		tokenB := (machine.A[0]*machine.Prize[1] - machine.A[1]*machine.Prize[0]) / (machine.A[0]*machine.B[1] - machine.A[1]*machine.B[0])

		if tokenA*machine.A[0]+tokenB*machine.B[0] == machine.Prize[0] && tokenA*machine.A[1]+tokenB*machine.B[1] == machine.Prize[1] {
			tokens += (tokenA * 3) + tokenB
		}
	}
	return tokens
}

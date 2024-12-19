package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	lines    []string
	progints []int
	prog     [][2]int
	output   []int
)

// Run executes the instructions using the given register A.
func run(regA int) {
	var regAVal, regB, regC, IP int
	output = []int{}
	regAVal = regA
	regB, _ = strconv.Atoi(lines[1][12:])
	regC, _ = strconv.Atoi(lines[2][12:])

	for IP < len(prog) {
		instr, operand := prog[IP][0], prog[IP][1]
		switch instr {
		case 0:
			regAVal /= 1 << getOperandValue(operand, regAVal, regB, regC)
		case 1:
			regB ^= operand
		case 2:
			regB = getOperandValue(operand, regAVal, regB, regC) % 8
		case 3:
			if regAVal != 0 {
				IP = operand - 1
			}
		case 4:
			regB ^= regC
		case 5:
			value := getOperandValue(operand, regAVal, regB, regC) % 8
			regB = value
			output = append(output, value)
		case 6:
			regB = regAVal / (1 << getOperandValue(operand, regAVal, regB, regC))
		case 7:
			regC = regAVal / (1 << getOperandValue(operand, regAVal, regB, regC))
		}
		IP++
	}
}

// getOperandValue returns the value based on the operand index.
func getOperandValue(operand, regA, regB, regC int) int {
	values := []int{operand, operand, operand, operand, regA, regB, regC}
	return values[operand]
}

func part1() string {
	run(parseRegisterValue(lines[0]))
	return strings.Join(intSliceToStringSlice(output), ",")
}

func part2(regA, posn int) int {
	if posn == len(progints) {
		return regA
	}
	for i := 0; i < 8; i++ {
		run(regA*8 + i)
		if len(output) > 0 && output[0] == progints[len(progints)-(posn+1)] {
			if result := part2(regA*8+i, posn+1); result != 0 {
				return result
			}
		}
	}
	return 0
}

func parseRegisterValue(line string) int {
	val, _ := strconv.Atoi(line[12:])
	return val
}

func intSliceToStringSlice(nums []int) []string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	return strs
}

func main() {

	data, _ := os.ReadFile("solutions/2024/Day17/input.txt")

	lines = strings.Split(strings.TrimSpace(string(data)), "\n")
	progintsStr := strings.Split(lines[4][9:], ",")
	for _, val := range progintsStr {
		num, _ := strconv.Atoi(val)
		progints = append(progints, num)
	}

	for i := 0; i < len(progints); i += 2 {
		prog = append(prog, [2]int{progints[i], progints[i+1]})
	}

	fmt.Println("Part 1:", part1())

	fmt.Println("Part 2:", part2(0, 0))
}

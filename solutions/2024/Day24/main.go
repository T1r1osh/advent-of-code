package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("solutions/2024/Day24/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	isInitValues := true
	values := map[string]bool{}
	operations := [][]string{}
	for scanner.Scan() {

		line := scanner.Text()
		if len(line) == 0 {
			isInitValues = false
			continue
		}

		if isInitValues {
			line := strings.Split(scanner.Text(), ": ")
			boolVal, _ := strconv.ParseBool(line[1])
			values[line[0]] = boolVal
		} else {
			line := strings.Split(scanner.Text(), " -> ")
			row := strings.Split(line[0], " ")
			row = append(row, line[1])
			operations = append(operations, row)
		}

	}

	log.Println(Task1(operations, values))
	log.Println(Task2(operations, values))

}

func Task2(operations [][]string, initValues map[string]bool) string {

	faulty := []string{}
	for _, operation := range operations {
		a := operation[0]
		gate := operation[1]
		b := operation[2]
		output := operation[3]

		// If the output of a gate is z, then the operation has to be XOR unless it is the last bit.
		if output[0] == 'z' && gate != "XOR" && output != "z45" {
			faulty = append(faulty, output)
			continue
		}

		// If the output of a gate is not z and the inputs are not x,y then it has to be AND / OR, but not XOR.
		if output[0] != 'z' && a[0] != 'x' && a[0] != 'y' && b[0] != 'x' && b[0] != 'y' && gate == "XOR" {
			faulty = append(faulty, output)
			continue
		}
		// If you have a XOR gate with inputs x, y, there must be another XOR gate with thihs gate as an input.
		// Search through all gates for an XOR-gate with this output as an input, if it does not exist, your original XOR gate is faulty.
		if gate == "XOR" && ((a[0] == 'x' && b[0] == 'y') || (a[0] == 'y' && b[0] == 'x')) &&
			a != "x00" && b != "x00" && a != "y00" && b != "y00" {
			isValid := false
			for _, line := range operations {
				if line[1] == "XOR" && (line[0] == output || line[2] == output) {
					isValid = true
					break
				}
			}
			if !isValid {
				faulty = append(faulty, output)
				continue
			}
		}

		// If you have an AND gate, there must be an OR gate with this output as an input
		// If that gate doesn't exist, the original AND gate is faulty.
		if gate == "AND" && ((a[0] == 'x' && b[0] == 'y') || (a[0] == 'y' && b[0] == 'x')) &&
			a != "x00" && b != "x00" && a != "y00" && b != "y00" {
			isValid := false
			for _, line := range operations {
				if line[1] == "OR" && (line[0] == output || line[2] == output) {
					isValid = true
					break
				}
			}
			if !isValid {
				faulty = append(faulty, output)
				continue
			}
		}
	}
	sort.Strings(faulty)
	return strings.Join(faulty, ",")
}

func Task1(operations [][]string, initValues map[string]bool) int {

	values := make(map[string]bool, len(initValues))

	for key, value := range initValues {
		values[key] = value
	}

	for len(operations) > 0 {

		current := operations[0]
		operations = operations[1:]

		if _, ok := values[current[0]]; !ok {
			operations = append(operations, current)
			continue
		}

		if _, ok := values[current[2]]; !ok {
			operations = append(operations, current)
			continue
		}
		bool1 := values[current[0]]
		bool2 := values[current[2]]
		switch current[1] {

		case "AND":
			values[current[3]] = bool1 && bool2
		case "OR":
			values[current[3]] = bool1 || bool2
		case "XOR":
			values[current[3]] = (bool1 || bool2) && !(bool1 && bool2)
		}
	}

	keys := make([]string, 0, len(values))

	for k := range values {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	binariNum := []bool{}
	for _, k := range keys {
		if k[:1] == "z" {
			binariNum = append(binariNum, values[k])
		}

	}
	return boolArrayToDecimal(binariNum)
}

func boolArrayToDecimal(arr []bool) int {
	result := 0
	n := len(arr)
	for i, val := range arr {
		if val {
			result += int(math.Pow(2, float64(n-i-1)))
		}
	}
	return result
}

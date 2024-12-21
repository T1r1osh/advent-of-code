package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var numPad map[string][2]int
var keyPad map[string][2]int

func main() {
	file, _ := os.Open("solutions/2024/Day21/input.txt")
	defer file.Close()

	codes := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		codes = append(codes, line)
	}

	log.Println(codes)

	numPad = make(map[string][2]int, 12)

	numPad["GAP"] = [2]int{0, 3}
	numPad["0"] = [2]int{1, 3}
	numPad["A"] = [2]int{2, 3}
	numPad["1"] = [2]int{0, 2}
	numPad["2"] = [2]int{1, 2}
	numPad["3"] = [2]int{2, 2}
	numPad["4"] = [2]int{0, 1}
	numPad["5"] = [2]int{1, 1}
	numPad["6"] = [2]int{2, 1}
	numPad["7"] = [2]int{0, 0}
	numPad["8"] = [2]int{1, 0}
	numPad["9"] = [2]int{2, 0}

	keyPad = make(map[string][2]int, 6)

	keyPad["GAP"] = [2]int{0, 0}
	keyPad["^"] = [2]int{1, 0}
	keyPad["A"] = [2]int{2, 0}
	keyPad["<"] = [2]int{0, 1}
	keyPad["v"] = [2]int{1, 1}
	keyPad[">"] = [2]int{2, 1}

	//	log.Println(numToKey("179A"))
	//log.Println(len(KeyToKey(KeyToKey(("^<<A^^A>>AvvvA")))))
	sum := 0
	for _, code := range codes {
		step1 := numToKey(code)
		step2 := KeyToKey(step1)
		step3 := KeyToKey(step2)
		log.Println(code)
		log.Println(step1)
		log.Println(step2)
		log.Println(step3)
		tmp := calculateComplexity(step3, code)
		log.Println(len(step3), code)
		sum += tmp
	}

	println(sum)
	//log.Println(KeyToKey("v<<A>>^A<A>AvA<^AA>A<vAAA>^A"))
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

func calculateComplexity(input, code string) int {

	length := len(input)
	codeStr := code[:3]
	codeNum, _ := strconv.Atoi(codeStr)
	return length * codeNum

}

func numToKey(code string) string {

	sb := strings.Builder{}

	codeAsArray := strings.Split(code, "")
	start := numPad["A"]
	for _, char := range codeAsArray {
		goal := numPad[char]
		distance := [2]int{goal[0] - start[0], goal[1] - start[1]}
		moveX := make([]string, abs(distance[0]))
		for i, _ := range moveX {
			key := ""
			if distance[0] < 0 {
				key = "<"
			} else {
				key = ">"
			}
			moveX[i] = key
		}

		moveY := make([]string, abs(distance[1]))
		for i, _ := range moveY {
			key := ""
			if distance[1] > 0 {
				key = "v"
			} else {
				key = "^"
			}
			moveY[i] = key
		}

		stepX := 0
		if distance[0] != 0 {
			stepX = (distance[0] / abs(distance[0]))
		}

		stepY := 0
		if distance[1] != 0 {
			stepY = (distance[1] / abs(distance[1]))
		}
		gapFound := false
		for i := 0; i < len(moveX); i++ {
			if start[0]+stepX*(i+1) == numPad["GAP"][0] && start[1] == numPad["GAP"][1] {
				gapFound = true
			}
		}
		if !gapFound {
			for _, move := range moveX {
				moveX = moveX[1:]
				start[0] += stepX
				sb.WriteString(move)
			}
		}

		for _, move := range moveY {
			moveY = moveY[1:]
			start[1] += stepY
			sb.WriteString(move)
		}

		for _, move := range moveX {
			start[0] += stepX
			sb.WriteString(move)
		}

		sb.WriteString("A")
		// x pos akkor < y negativ akkor ^ x neg > ha x pos v
	}
	return sb.String()
}

//+---+---+---+
//|   | ^ | A |
//+---+---+---+
//| < | v | > |
//+---+---+---+

func KeyToKey(code string) string {

	sb := strings.Builder{}

	codeAsArray := strings.Split(code, "")
	start := keyPad["A"]
	for _, char := range codeAsArray {
		goal := keyPad[char]
		distance := [2]int{goal[0] - start[0], goal[1] - start[1]}
		moveX := make([]string, abs(distance[0]))
		for i, _ := range moveX {
			key := ""
			if distance[0] < 0 {
				key = "<"
			} else {
				key = ">"
			}
			moveX[i] = key
		}

		moveY := make([]string, abs(distance[1]))
		for i, _ := range moveY {
			key := ""
			if distance[1] > 0 {
				key = "v"
			} else {
				key = "^"
			}
			moveY[i] = key
		}

		stepX := 0
		if distance[0] != 0 {
			stepX = (distance[0] / abs(distance[0]))
		}

		stepY := 0
		if distance[1] != 0 {
			stepY = (distance[1] / abs(distance[1]))
		}
		gapFound := false
		for i := 0; i < len(moveX); i++ {
			if start[0]+stepX*(i+1) == numPad["GAP"][0] && start[1] == numPad["GAP"][1] {
				gapFound = true
			}
		}
		if !gapFound {
			for _, move := range moveX {
				moveX = moveX[1:]
				start[0] += stepX
				sb.WriteString(move)
			}
		}
		for _, move := range moveY {
			moveY = moveY[1:]
			start[1] += stepY
			sb.WriteString(move)
		}

		for _, move := range moveX {
			start[0] += stepX
			sb.WriteString(move)
		}

		sb.WriteString("A")
		// x pos akkor < y negativ akkor ^ x neg > ha x pos v
	}
	return sb.String()
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

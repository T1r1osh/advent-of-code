package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var directions = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {

	file, _ := os.Open("solutions/2024/Day15/input.txt")
	defer file.Close()

	warehouse := [][]string{}
	scanner := bufio.NewScanner(file)
	rowindex := 0
	robotPos := [2]int{}
	sequence := []string{}
	isCommands := false
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		if len(line) == 0 {
			isCommands = true
			continue
		}

		if !isCommands {
			warehouse = append(warehouse, []string{})
			warehouse[rowindex] = make([]string, len(line))
			for i, col := range line {
				warehouse[rowindex][i] = col
				if col == "@" {
					robotPos[0] = i
					robotPos[1] = rowindex
				}
			}
		} else {
			sequence = append(sequence, line...)
		}
		rowindex++
	}

	log.Println(Task1(warehouse, sequence, robotPos))
	log.Println(Task2(warehouse, sequence))
}

func Task1(warehouse [][]string, sequence []string, robotPos [2]int) int {

	copyWarehouse := make([][]string, len(warehouse))
	for i, row := range warehouse {
		copyWarehouse[i] = make([]string, len(row))
		copy(copyWarehouse[i], row)
	}

	for _, cmd := range sequence {
		copyWarehouse, robotPos = processStep(robotPos, copyWarehouse, cmd)
	}

	sum := 0
	for y, row := range copyWarehouse {
		for x, col := range row {
			if col == "O" {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func Task2(warehouse [][]string, sequence []string) int {

	copyWarehouse := make([][]string, len(warehouse))
	for i, row := range warehouse {
		copyWarehouse[i] = make([]string, len(row))
		copy(copyWarehouse[i], row)
	}
	expandedWarehouse, robotPos := expandGrid(copyWarehouse)
	//printMap(expandedWarehouse)
	for _, cmd := range sequence {
		expandedWarehouse, robotPos = processStep(robotPos, expandedWarehouse, cmd)
	}

	sum := 0
	for y, row := range expandedWarehouse {
		for x, col := range row {
			if col == "O" || col == "[" {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func processStep(robotPos [2]int, warehouse [][]string, cmd string) ([][]string, [2]int) {
	var direction [2]int
	switch cmd {
	case "^":
		direction = directions[0]
	case ">":
		direction = directions[1]
	case "v":
		direction = directions[2]
	case "<":
		direction = directions[3]
	}

	newX, newY := robotPos[0]+direction[0], robotPos[1]+direction[1]

	if !isValidStep([2]int{newX, newY}, warehouse) {
		return warehouse, robotPos
	}

	switch warehouse[newY][newX] {
	case ".":
		warehouse[robotPos[1]][robotPos[0]] = "."
		warehouse[newY][newX] = "@"
		return warehouse, [2]int{newX, newY}
	case "#":
		return warehouse, robotPos
	case "O":
		for isValidStep([2]int{newX, newY}, warehouse) {
			switch warehouse[newY][newX] {
			case ".":
				warehouse[robotPos[1]][robotPos[0]] = "."
				warehouse[newY][newX] = "O"
				warehouse[robotPos[1]+direction[1]][robotPos[0]+direction[0]] = "@"
				return warehouse, [2]int{robotPos[0] + direction[0], robotPos[1] + direction[1]}
			case "#":
				return warehouse, robotPos
			}
			newX += direction[0]
			newY += direction[1]
		}
	case "[", "]":
		isPossible := moveBoxes([2]int{newX, newY}, direction, warehouse)
		if !isPossible {
			return warehouse, robotPos
		}
		warehouse[newY][newX] = "@"
		warehouse[robotPos[1]][robotPos[0]] = "."
		return warehouse, [2]int{newX, newY}
	}
	return warehouse, robotPos
}

func moveBoxes(start, dir [2]int, warehouse [][]string) bool {
	if dir[1] == 0 && (dir[0] == 1 || dir[0] == -1) {
		return moveBoxesHorizontally(start, dir, warehouse)
	}
	if dir[0] == 0 && (dir[1] == 1 || dir[1] == -1) {
		return moveBoxesVertically(start, dir, warehouse)
	}
	return false
}

func moveBoxesHorizontally(start, dir [2]int, warehouse [][]string) bool {
	newX, newY := start[0]+dir[0], start[1]+dir[1]
	switch warehouse[newY][newX] {
	case "#":
		return false
	case ".":
		warehouse[newY][newX], warehouse[start[1]][start[0]] = warehouse[start[1]][start[0]], warehouse[newY][newX]
		return true
	case "]", "[":
		isPossible := moveBoxes([2]int{newX, newY}, dir, warehouse)
		if !isPossible {
			return false
		}
		warehouse[start[1]][start[0]], warehouse[newY][newX] = warehouse[newY][newX], warehouse[start[1]][start[0]]
	}
	return true
}

func moveBoxesVertically(start, dir [2]int, warehouse [][]string) bool {

	next := [][2]int{{start[0], start[1]}}
	if warehouse[start[1]][start[0]] == "]" {
		next = append(next, [2]int{start[0] - 1, start[1]})
	} else {
		next = append(next, [2]int{start[0] + 1, start[1]})
	}

	visited := make(map[[2]int]struct{})
	visitedSlice := [][2]int{}

	for len(next) != 0 {
		process := next[0]
		next = next[1:]

		if _, ok := visited[process]; ok {
			continue
		}
		visited[process] = struct{}{}
		visitedSlice = append(visitedSlice, process)

		newX, newY := process[0]+dir[0], process[1]+dir[1]
		switch warehouse[newY][newX] {
		case ".":
			continue
		case "#":
			return false
		case "]":
			next = append(next, [2]int{newX, newY})
			next = append(next, [2]int{newX - 1, newY})
		case "[":
			next = append(next, [2]int{newX, newY})
			next = append(next, [2]int{newX + 1, newY})
		}
	}

	for i := len(visitedSlice) - 1; i >= 0; i-- {
		x, y := visitedSlice[i][0]+dir[0], visitedSlice[i][1]+dir[1]
		warehouse[y][x] = warehouse[visitedSlice[i][1]][visitedSlice[i][0]]
		warehouse[visitedSlice[i][1]][visitedSlice[i][0]] = "."
	}
	return true
}

func isValidStep(pos [2]int, warehouse [][]string) bool {
	if pos[0] < 0 || pos[1] < 0 || pos[0] >= len(warehouse[0]) || pos[1] >= len(warehouse) {
		return false
	}
	return true
}

func printMap(warehouse [][]string) {
	for _, row := range warehouse {
		for _, col := range row {
			print(col)
		}
		println()
	}
}

func expandGrid(warehouse [][]string) ([][]string, [2]int) {
	//printMap(warehouse)
	bigWarehouse := [][]string{}
	robotPos := [2]int{}
	for j, row := range warehouse {
		finalRow := []string{}
		for i, col := range row {
			switch col {
			case ".":
				finalRow = append(finalRow, []string{".", "."}...)
			case "#":
				finalRow = append(finalRow, []string{"#", "#"}...)
			case "O":
				finalRow = append(finalRow, []string{"[", "]"}...)
			case "@":
				robotPos[0] = i * 2
				robotPos[1] = j
				finalRow = append(finalRow, []string{"@", "."}...)
			}
		}
		bigWarehouse = append(bigWarehouse, finalRow)
	}
	//printMap(bigWarehouse)
	return bigWarehouse, robotPos
}

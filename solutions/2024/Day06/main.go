package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

type Point struct {
	Y int
	X int
}

func (p *Point) Add(p2 Point) {
	p.X = p.X + p2.X
	p.Y = p.Y + p2.Y
}

type Dir int

const (
	Up Dir = iota
	Right
	Down
	Left
)

var (
	Move = []Point{
		{X: 0, Y: -1}, // Up 0
		{X: 1, Y: 0},  // Right 1
		{X: 0, Y: 1},  // Down 2
		{X: -1, Y: 0}} // Left 3
)

func (d *Dir) ChangeDir() {
	switch *d {
	case Up:
		*d = Right
		return
	case Right:
		*d = Down
		return
	case Down:
		*d = Left
		return
	case Left:
		*d = Up
		return
	}
}

func CanMove(guardPos Point, dir *Dir, inputMap [][]string) (canmove bool, exit bool) {

	maxY := len(inputMap) - 1
	maxX := len(inputMap[0]) - 1

	if *dir == Up {
		if guardPos.Y-1 < 0 {
			canmove = false
			exit = true
			return
		}
		canmove = inputMap[guardPos.Y-1][guardPos.X] != "#"
		exit = false
		return
	}
	if *dir == Right {
		if guardPos.X+1 > maxX {
			canmove = false
			exit = true
			return
		}
		canmove = inputMap[guardPos.Y][guardPos.X+1] != "#"
		exit = false
		return
	}
	if *dir == Down {
		if guardPos.Y+1 > maxY {
			canmove = false
			exit = true
			return
		}
		canmove = inputMap[guardPos.Y+1][guardPos.X] != "#"
		exit = false
		return
	}
	if *dir == Left {
		if guardPos.X-1 < 0 {
			canmove = false
			exit = true
			return
		}
		canmove = inputMap[guardPos.Y][guardPos.X-1] != "#"
		exit = false
		return
	}
	panic("invalid dir")
}

func Task1(guardPos Point, inputMap [][]string) (map[Point]bool, bool) {

	infinite := false
	dir := Up
	positions := map[Point]bool{}

	dirMappedPositions := map[Dir]map[Point]int{}
	for _, direction := range []Dir{Up, Right, Down, Left} {
		dirMappedPositions[direction] = make(map[Point]int)
	}
	dirMappedPositions[dir][guardPos] = 1
	for {
		positions[guardPos] = true
		canmove, exit := CanMove(guardPos, &dir, inputMap)
		if exit {
			break
		}
		if !canmove {
			dir.ChangeDir()
			continue
		}
		guardPos.Add(Move[dir])
		dirMappedPositions[dir][guardPos] = dirMappedPositions[dir][guardPos] + 1

		if dirMappedPositions[dir][guardPos] > 1 {
			// guard is in a loop
			infinite = true
			break
		}

	}
	return positions, infinite
}

func Task2(guardPos Point, inputMap [][]string) int {

	positions, _ := Task1(guardPos, inputMap)
	positions[guardPos] = false
	goodBlockade := 0

	for pos := range positions {
		if positions[pos] {
			newMap := ChangeMap(inputMap, pos)
			_, loop := Task1(guardPos, newMap)
			if loop {
				goodBlockade++
			}
		}
	}
	return goodBlockade
}

func ChangeMap(inputMap [][]string, p Point) (changedMap [][]string) {
	changedMap = make([][]string, len(inputMap))
	for i := range inputMap {
		changedMap[i] = make([]string, len(inputMap[0]))
		copy(changedMap[i], inputMap[i])
	}
	changedMap[p.Y][p.X] = "#"
	return
}

func main() {
	file, _ := os.Open("solutions/2024/Day06/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	inputMap := [][]string{}
	rownr := 0
	guardPos := Point{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		inputMap = append(inputMap, line)
		guardX := slices.Index(line, "^")
		if guardX != -1 {
			guardPos = Point{Y: rownr, X: guardX}
		}
		rownr++
	}
	answer1, _ := Task1(guardPos, inputMap)
	log.Println(len(answer1))
	log.Println(Task2(guardPos, inputMap))
}

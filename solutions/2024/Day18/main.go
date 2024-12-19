package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const MAZEDIM = 71

func main() {
	file, _ := os.Open("solutions/2024/Day18/input.txt")
	defer file.Close()

	maze := make([][]string, MAZEDIM)
	for i := 0; i < len(maze); i++ {
		maze[i] = make([]string, MAZEDIM)
	}

	scanner := bufio.NewScanner(file)
	corruptedBytes := [][2]int{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		y, _ := strconv.Atoi(line[1])
		x, _ := strconv.Atoi(line[0])
		newLine := [2]int{x, y}
		corruptedBytes = append(corruptedBytes, newLine)
	}

	for i := 0; i < 1024; i++ {
		maze[corruptedBytes[i][1]][corruptedBytes[i][0]] = "#"
	}

	log.Println("Task1", Task1(maze))

	for j := 1024; j < len(corruptedBytes); j++ {
		maze[corruptedBytes[j][1]][corruptedBytes[j][0]] = "#"
		if Task1(maze) == -1 {
			log.Println("Task2: ", corruptedBytes[j][0], ",", corruptedBytes[j][1])
			break
		}
	}

}

func Task1(maze [][]string) int {
	queToExp := [][3]int{}
	visited := map[[2]int]bool{}
	visited[[2]int{0, 0}] = true

	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queToExp = append(queToExp, [3]int{0, 0, 0})
	ctr := 0
	for len(queToExp) > 0 {
		ctr++
		x := queToExp[0][0]
		y := queToExp[0][1]
		steps := queToExp[0][2]
		queToExp = append(queToExp[:0], queToExp[1:]...)

		if x == MAZEDIM-1 && y == MAZEDIM-1 {
			return steps
		}

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			if (0 <= nx && nx < MAZEDIM) && (0 <= ny && ny < MAZEDIM) && maze[ny][nx] != "#" && !visited[[2]int{ny, nx}] {
				queToExp = append(queToExp, [3]int{nx, ny, steps + 1})
				visited[[2]int{ny, nx}] = true
			}
		}
	}
	return -1
}

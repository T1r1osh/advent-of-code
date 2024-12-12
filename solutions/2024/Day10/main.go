package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("solutions/2024/Day10/input.txt")
	defer file.Close()

	valMap := map[int]map[int]int{}
	grid := [][]int{}
	loop := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		intline := make([]int, len(line))
		valMap[loop] = make(map[int]int, 0)
		for i, val := range line {
			intline[i], _ = strconv.Atoi(val)
			valMap[loop][i] = intline[i]
		}
		grid = append(grid, intline)
		loop++
	}
	log.Println(Task1(grid))
	log.Println(Task2(grid))
}

func Task1(grid [][]int) int {
	trailCountSum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// trail start
			if grid[i][j] == 0 {
				visited := make([][]bool, len(grid))
				for p := 0; p < len(visited); p++ {
					visited[p] = make([]bool, len(grid[p]))
				}

				// mark the nines with a good trail
				markVisited(grid, i, j, visited)

				for p := 0; p < len(visited); p++ {
					for q := 0; q < len(visited[p]); q++ {
						if visited[p][q] {
							trailCountSum++
						}
					}
				}
			}
		}
	}
	return trailCountSum
}

func Task2(grid [][]int) int {
	trailCountSum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				trailCountSum += getDistincTrailCount(grid, i, j)
			}
		}
	}
	return trailCountSum
}

func markVisited(grid [][]int, i, j int, visited [][]bool) {

	if grid[i][j] == 9 {
		// trail end
		visited[i][j] = true
		return
	}
	for _, neighbour := range getNeighbors(i, j) {
		ni, nj := neighbour[0], neighbour[1]
		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) && grid[ni][nj]-grid[i][j] == 1 {
			markVisited(grid, ni, nj, visited)
		}
	}
}

func getNeighbors(i, j int) [4][2]int {
	return [4][2]int{
		{i - 1, j},
		{i + 1, j},
		{i, j - 1},
		{i, j + 1},
	}
}

func getDistincTrailCount(grid [][]int, i, j int) int {
	if grid[i][j] == 9 {
		return 1
	}
	count := 0
	for _, neighbour := range getNeighbors(i, j) {
		ni, nj := neighbour[0], neighbour[1]
		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) && grid[ni][nj]-grid[i][j] == 1 {
			count += getDistincTrailCount(grid, ni, nj)
		}
	}
	return count
}

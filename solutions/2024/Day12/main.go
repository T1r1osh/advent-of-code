package main

import (
	"log"
	"os"
	"sort"
	"strings"
)

type Point struct {
	Y, X int
}

func main() {

	content, _ := os.ReadFile("solutions/2024/Day12/input.txt")
	farm := parseInput(string(content))

	log.Println(Task1(farm))
	log.Println(Task2(farm))

}

func parseInput(input string) [][]byte {
	lines := strings.Split(input, "\n")
	farm := make([][]byte, len(lines))
	for i, line := range lines {
		farm[i] = []byte(line)
	}
	return farm
}

func dfs(farm [][]byte, visited [][]bool, i, j int, sides *[][3]int) (int, int) {

	visited[i][j] = true
	neighbors := getNeighbors(i, j)
	perimeter := 0
	area := 1
	for _, n := range neighbors {
		ni, nj := n[0], n[1]
		if ni < 0 || ni >= len(farm) || nj < 0 || nj >= len(farm[ni]) || farm[ni][nj] != farm[i][j] {
			if sides != nil {
				*sides = append(*sides, [3]int{ni, nj, n[2]})
			}
			perimeter++
		} else if !visited[ni][nj] {
			a, p := dfs(farm, visited, ni, nj, sides)
			area += a
			perimeter += p
		}
	}
	return area, perimeter
}

func getNeighbors(i, j int) [][3]int {
	return [][3]int{
		{i - 1, j, 0},
		{i + 1, j, 1},
		{i, j - 1, 2},
		{i, j + 1, 3},
	}
}
func Task1(farm [][]byte) int {
	visited := make([][]bool, len(farm))
	for i := range visited {
		visited[i] = make([]bool, len(farm[i]))
	}
	sum := 0
	for i := 0; i < len(farm); i++ {
		for j := 0; j < len(farm[i]); j++ {
			if !visited[i][j] {
				area, perimeter := dfs(farm, visited, i, j, nil)
				sum += area * perimeter
			}
		}
	}
	return sum
}

func getSideCount(sides [][3]int) int {
	sideMap := make(map[[3]int]bool)

	sort.Slice(sides, func(i, j int) bool {
		if sides[i][0] == sides[j][0] {
			return sides[i][1] < sides[j][1]
		}
		return sides[i][0] < sides[j][0]
	})

	sideCount := 0

	for _, s := range sides {
		getCombinations := getNeighbors(s[0], s[1])
		combFound := false

		for _, c := range getCombinations {
			c[2] = s[2]
			if _, found := sideMap[c]; found {
				combFound = true
			}
		}
		if !combFound {
			sideCount++
		}

		sideMap[s] = true

	}

	return sideCount
}

func Task2(farm [][]byte) int {
	visited := make([][]bool, len(farm))
	for i := range visited {
		visited[i] = make([]bool, len(farm[i]))
	}
	sum := 0
	for i := 0; i < len(farm); i++ {
		for j := 0; j < len(farm[i]); j++ {
			if !visited[i][j] {
				sides := make([][3]int, 0)
				area, _ := dfs(farm, visited, i, j, &sides)
				sideCount := getSideCount(sides)
				sum += area * sideCount
			}
		}
	}
	return sum
}

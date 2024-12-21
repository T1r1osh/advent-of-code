package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// change this for the example
const TIME_SAVE = 101

func main() {
	file, _ := os.Open("solutions/2024/Day20/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	maze := [][]string{}
	start := [2]int{}
	finish := [2]int{}
	rowindex := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		maze = append(maze, line)
		for i, col := range line {
			if col == "S" {
				start[0] = i
				start[1] = rowindex
			}
			if col == "E" {
				finish[0] = i
				finish[1] = rowindex
				maze[rowindex][i] = "."
			}
		}
		rowindex++
	}

	pos_x := start[0]
	pos_y := start[1]

	traversal_list := [][2]int{}
	finished := false
	// get the only path
	for !finished {
		coordinates := [2]int{pos_x, pos_y}
		traversal_list = append(traversal_list, coordinates)
		if pos_x == finish[0] && pos_y == finish[1] {
			finished = true
			break
		}
		maze[pos_y][pos_x] = "X"
		if maze[pos_y][pos_x+1] == "." {
			pos_x = pos_x + 1
		} else if maze[pos_y][pos_x-1] == "." {
			pos_x = pos_x - 1
		} else if maze[pos_y+1][pos_x] == "." {
			pos_y = pos_y + 1
		} else if maze[pos_y-1][pos_x] == "." {
			pos_y = pos_y - 1
		}
	}

	oneSecCheat := 0
	twentySecCheat := 0
	for a := 0; a < len(traversal_list); a++ {
		for b := a + TIME_SAVE; b < len(traversal_list); b++ {

			distance := [2]int{}
			distance[0] = traversal_list[a][0] - traversal_list[b][0]
			distance[1] = traversal_list[a][1] - traversal_list[b][1]

			if distance[0] < 0 {
				distance[0] = distance[0] * -1
			}
			if distance[1] < 0 {
				distance[1] = distance[1] * -1
			}
			total_distance := distance[0] + distance[1]
			if total_distance <= 20 {
				if b-a-total_distance >= TIME_SAVE-1 {
					twentySecCheat++
				}
			}
			if total_distance <= 2 {
				oneSecCheat++
			}
		}
	}

	log.Println("Task1 : ", oneSecCheat)
	log.Println("Task2 : ", twentySecCheat)
}

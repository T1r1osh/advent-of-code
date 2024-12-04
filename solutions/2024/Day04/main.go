package main

import (
	"bufio"
	"os"
	"strings"
)

func Task1(matrix [][]string) int {
	counter := 0
	for y, row := range matrix {
		for x, val := range row {
			if val == "X" {
				//search right
				if x < len(row)-3 {
					if strings.Join([]string{val, matrix[y][x+1], matrix[y][x+2], matrix[y][x+3]}, "") == "XMAS" {
						counter++
					}
				}
				//search left
				if x > 2 {
					if strings.Join([]string{val, matrix[y][x-1], matrix[y][x-2], matrix[y][x-3]}, "") == "XMAS" {
						counter++
					}
				}
				//search up
				if y > 2 {
					if strings.Join([]string{val, matrix[y-1][x], matrix[y-2][x], matrix[y-3][x]}, "") == "XMAS" {
						counter++
					}
				}
				//search down
				if y < len(matrix)-3 {
					if strings.Join([]string{val, matrix[y+1][x], matrix[y+2][x], matrix[y+3][x]}, "") == "XMAS" {
						counter++
					}
				}
				//search topL downR
				if x < len(row)-3 && y < len(matrix)-3 {
					if strings.Join([]string{val, matrix[y+1][x+1], matrix[y+2][x+2], matrix[y+3][x+3]}, "") == "XMAS" {
						counter++
					}
				}
				//search downL upR
				if x < len(row)-3 && y > 2 {
					if strings.Join([]string{val, matrix[y-1][x+1], matrix[y-2][x+2], matrix[y-3][x+3]}, "") == "XMAS" {
						counter++
					}
				}
				//search rightT downL
				if x > 2 && y < len(matrix)-3 {
					if strings.Join([]string{val, matrix[y+1][x-1], matrix[y+2][x-2], matrix[y+3][x-3]}, "") == "XMAS" {
						counter++
					}
				}
				//search downR upL
				if x > 2 && y > 2 {
					if strings.Join([]string{val, matrix[y-1][x-1], matrix[y-2][x-2], matrix[y-3][x-3]}, "") == "XMAS" {
						counter++
					}
				}
			}
		}
	}

	return counter

}

func Task2(matrix [][]string) int {
	cross := 0
	for y, row := range matrix {
		for x, val := range row {
			if val == "A" && (y > 0 && x > 0 && x < len(matrix)-1 && y < len(matrix)-1) {
				word1 := strings.Join([]string{matrix[y-1][x-1], val, matrix[y+1][x+1]}, "")
				word2 := strings.Join([]string{matrix[y+1][x-1], val, matrix[y-1][x+1]}, "")
				if (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
					cross++
				}
			}
		}
	}
	return cross
}

func main() {

	file, _ := os.Open("solutions/2024/Day04/input.txt")

	scanner := bufio.NewScanner(file)
	matrix := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}

	task1 := Task1(matrix)
	task2 := Task2(matrix)

	println("Task 1 answer is: ", task1)
	println("Task 2 answer is: ", task2)

}

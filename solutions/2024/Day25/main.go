package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("solutions/2024/Day25/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	keys := map[int][5]int{}
	locks := map[int][5]int{}

	isKey := false
	firstRow := true
	keyIndex := 0
	lockIndex := 0
	for scanner.Scan() {

		line := scanner.Text()
		if len(line) == 0 {
			if isKey {
				keyIndex++
			} else {
				lockIndex++
			}
			firstRow = true
			isKey = false
			continue
		}

		if firstRow && line == "....." {
			isKey = true
			firstRow = false
		} else if firstRow && line == "#####" {
			isKey = false
			firstRow = false
		}

		if isKey {
			if _, ok := keys[keyIndex]; !ok {
				keys[keyIndex] = [5]int{}
			}
			cols := strings.Split(line, "")
			for i, col := range cols {
				if col == "#" {
					arr := keys[keyIndex]
					arr[i]++
					keys[keyIndex] = arr
				}
			}
		} else {
			if _, ok := locks[lockIndex]; !ok {
				locks[lockIndex] = [5]int{}
			}
			cols := strings.Split(line, "")
			for i, col := range cols {
				if col == "#" {
					arr := locks[lockIndex]
					arr[i]++
					locks[lockIndex] = arr
				}
			}
		}

	}
	sum := 0
	for _, key := range keys {
		for _, lock := range locks {
			good := true
			for i := range lock {
				if key[i]+lock[i] > 7 {
					good = false
				}
			}
			if good {
				sum++
			}
		}
	}
	println(sum)
}

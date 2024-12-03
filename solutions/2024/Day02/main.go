package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Floor struct {
	levels []int
}

func isSafeFloor(floor []int) bool {

	newfloor := Floor{floor}

	return newfloor.isSafe()
}

func isSafeFloorWithDampener(floor []int) bool {

	newfloor := Floor{floor}

	return newfloor.isSafeWithDampener()
}

func (f *Floor) isSafeAsc() bool {
	for i := 0; i < len(f.levels)-1; i++ {
		diff := f.levels[i] - f.levels[i+1]
		if -3 > diff || diff > -1 {
			return false
		}
	}
	return true
}

func (f *Floor) isSafeDesc() bool {
	for i := 0; i < len(f.levels)-1; i++ {
		diff := f.levels[i] - f.levels[i+1]
		if 1 > diff || diff > 3 {
			return false
		}
	}
	return true
}

func (f *Floor) isSafeWithDampener() bool {
	if f.isSafe() {
		return true
	}

	for i := 0; i < len(f.levels); i++ {
		dampedLevels := make([]int, 0, len(f.levels)-1)
		dampedLevels = append(dampedLevels, f.levels[:i]...)
		dampedLevels = append(dampedLevels, f.levels[i+1:]...)
		temp := Floor{dampedLevels}
		if temp.isSafe() {
			return true
		}
	}

	return false
}

func (f *Floor) isSafe() bool {
	return f.isSafeAsc() || f.isSafeDesc()
}

func main() {

	file, err := os.Open("solutions/2024/Day02/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safereports := 0
	safeDamper := 0
	for scanner.Scan() {
		line := scanner.Text()
		floorStr := strings.Split(line, " ")
		floorInt := make([]int, len(floorStr))
		for index, item := range floorStr {
			//we are assume that every input is valid
			floorInt[index], _ = strconv.Atoi(item)

		}
		if isSafeFloor(floorInt) {
			safereports++
		}
		if isSafeFloorWithDampener(floorInt) {
			safeDamper++
		}
	}

	println("Safe reports:", safereports)
	println("Safe reports with dampener:", safeDamper)

}

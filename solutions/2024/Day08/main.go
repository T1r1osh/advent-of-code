package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Point struct {
	Y, X int
}

func GetDoubleSlope(p1, p2 Point) Point {
	return Point{(p2.Y - p1.Y) * 2, (p2.X - p1.X) * 2}
}

func GetSlope(p1, p2 Point) Point {
	return Point{(p2.Y - p1.Y), (p2.X - p1.X)}
}

func CheckAntinodePos(p1, last Point) bool {
	// is Inside?
	return p1.X >= 0 && p1.X < last.X && p1.Y >= 0 && p1.Y < last.Y
}

func calculateAntinodePos(p1, p2 Point) []Point {
	slope := GetDoubleSlope(p1, p2)
	antinodes := make([]Point, 2)
	antinodes[0] = Point{p2.Y - slope.Y, p2.X - slope.X}
	antinodes[1] = Point{p1.Y + slope.Y, p1.X + slope.X}
	return antinodes
}

func calculateAntinodePosInline(p1, p2, lastCoord Point) []Point {

	antinodes := []Point{}
	slope := GetSlope(p1, p2)

	newPoint := Point{p1.Y, p1.X}
	for {
		if newPoint.X+slope.X >= lastCoord.X || newPoint.X < 0 {
			break
		}
		newPoint = Point{newPoint.Y + slope.Y, newPoint.X + slope.X}
		antinodes = append(antinodes, newPoint)
	}

	newPoint2 := Point{p2.Y, p2.X}
	for {
		if newPoint2.Y-slope.Y >= lastCoord.Y || newPoint2.Y-slope.Y < 0 {
			break
		}
		newPoint2 = Point{newPoint2.Y - slope.Y, newPoint2.X - slope.X}
		antinodes = append(antinodes, newPoint2)
	}

	return antinodes
}

func GetAntinodePositions(antennaLocation map[string][]Point, lastCoord Point, inline bool) map[Point]bool {
	antinodePos := map[Point]bool{}
	for _, freq := range antennaLocation {
		for j, point := range freq {
			antinodes := []Point{}
			for k := j + 1; k < len(freq); k++ {
				if inline {
					antinodes = append(antinodes, calculateAntinodePosInline(point, freq[k], lastCoord)...)
				} else {
					antinodes = append(antinodes, calculateAntinodePos(point, freq[k])...)
				}
			}
			for _, a1 := range antinodes {
				if CheckAntinodePos(a1, lastCoord) {
					antinodePos[a1] = true
				}
			}
		}
	}
	return antinodePos
}

func main() {

	file, _ := os.Open("solutions/2024/Day08/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	antennaMap := [][]string{}
	antennaLocation := map[string][]Point{}
	rownr := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		antennaMap = append(antennaMap, line)

		for i, val := range line {
			if val != "." {
				antennaLocation[val] = append(antennaLocation[val], Point{Y: rownr, X: i})
			}
		}
		rownr++
	}

	//Task1
	log.Println(len(GetAntinodePositions(antennaLocation, Point{len(antennaMap), len(antennaMap)}, false)))
	//Task2
	log.Println(len(GetAntinodePositions(antennaLocation, Point{len(antennaMap), len(antennaMap)}, true)))
}

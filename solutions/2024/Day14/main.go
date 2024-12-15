package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	WIDTH  = 101
	HEIGHT = 103
)

type Guard struct {
	pos, velocity [2]int
}

func main() {

	file, _ := os.Open("solutions/2024/Day14/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	guards := []Guard{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		p := parseCoordinates(parts[0][2:]) // Remove "p=" prefix
		v := parseCoordinates(parts[1][2:]) // Remove "v=" prefix
		guards = append(guards, Guard{pos: p, velocity: v})
	}

	log.Println(Task1(guards))
	log.Println(Task2(guards))

}

func Task1(guards []Guard) int {

	guardMap := map[string]int{}
	guardMap["TL"] = 0
	guardMap["TR"] = 0
	guardMap["BL"] = 0
	guardMap["BR"] = 0
	for _, guard := range guards {
		finalPos := Move(guard, 100)
		// top left
		if finalPos[0] < WIDTH/2 && finalPos[1] < HEIGHT/2 {
			guardMap["TL"]++
		}
		// top right
		if finalPos[0] > WIDTH/2 && finalPos[1] < HEIGHT/2 {
			guardMap["TR"]++
		}
		// bot left
		if finalPos[0] < WIDTH/2 && finalPos[1] > HEIGHT/2 {
			guardMap["BL"]++
		}
		// bot right
		if finalPos[0] > WIDTH/2 && finalPos[1] > HEIGHT/2 {
			guardMap["BR"]++
		}
	}
	safetyFactor := 1
	for _, quadrant := range guardMap {
		safetyFactor = safetyFactor * quadrant
	}
	return safetyFactor
}

func Task2(guards []Guard) int {

	seconds := 0
	for {
		seconds++
		canvas := make([][]int, HEIGHT)
		for i := 0; i < HEIGHT; i++ {
			canvas[i] = make([]int, WIDTH)
		}

		bad := false
		for _, guard := range guards {
			finalPos := Move(guard, seconds)
			canvas[finalPos[1]][finalPos[0]]++
			if canvas[finalPos[1]][finalPos[0]] > 1 {
				bad = true
			}

		}
		if !bad {
			/* fileName := "solutions/2024/Day14/canvas_data.txt"
			// Open the file for writing
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return 0
			}
			defer file.Close()

			// Write map data to the file
			for i := 0; i < len(canvas); i++ {
				line := fmt.Sprintf("%v\n", canvas[i])
				_, err := file.WriteString(line)
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return 0
				}
			} */
			break
		}
	}

	return seconds
}

func Move(guard Guard, seconds int) [2]int {
	px, py := guard.pos[0], guard.pos[1]
	vx, vy := guard.velocity[0], guard.velocity[1]

	// Compute new positions
	nx := (px + seconds*vx) % WIDTH
	ny := (py + seconds*vy) % HEIGHT

	// Handle wrap-around for negative values
	if nx < 0 {
		nx += WIDTH
	}
	if ny < 0 {
		ny += HEIGHT
	}
	return [2]int{nx, ny}
}

func parseCoordinates(coord string) [2]int {

	parts := strings.Split(coord, ",")

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return [2]int{x, y}
}

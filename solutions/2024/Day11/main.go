package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("solutions/2024/Day11/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stones := strings.Split(scanner.Text(), " ")
	stonesInt := make([]int, len(stones))
	for i, stone := range stones {
		stonesInt[i], _ = strconv.Atoi(stone)
	}

	log.Println(Task1(stonesInt, 25))
	log.Println(Task1(stonesInt, 75))
}

func performOP(num int) (int, int) {

	if num == 0 {
		return 1, -1
	}

	numStr := fmt.Sprintf("%d", num)

	if len(numStr)%2 == 0 { // even
		newNumLeftStr := numStr[:len(numStr)/2]
		newNumRightStr := numStr[len(numStr)/2:]
		newNumLeft, _ := strconv.Atoi(newNumLeftStr)
		newNumRight, _ := strconv.Atoi(newNumRightStr)

		return newNumLeft, newNumRight
	}
	return num * 2024, -1
}

func Task1(numArr []int, blinks int) int {

	numFreq := make(map[int]int)

	for _, num := range numArr {
		if _, found := numFreq[num]; found {
			numFreq[num]++
		} else {
			numFreq[num] = 1
		}

	}

	for i := 1; i <= blinks; i++ {
		//fmt.Println(numFreq)
		newFreq := make(map[int]int)
		for num, freq := range numFreq {

			num1, num2 := performOP(num)
			if _, found := newFreq[num1]; found {
				newFreq[num1] += freq
			} else {
				newFreq[num1] = freq
			}

			if num2 != -1 {
				if _, found := newFreq[num2]; found {
					newFreq[num2] += freq
				} else {
					newFreq[num2] = freq
				}
			}
		}

		numFreq = newFreq
	}

	countStones := 0
	for _, freq := range numFreq {
		countStones += freq
	}

	return countStones

}

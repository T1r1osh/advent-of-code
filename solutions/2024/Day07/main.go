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

	file, _ := os.Open("solutions/2024/Day07/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	task1sum := 0
	task2sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		goalStr := strings.Split(line, ": ")
		numberStr := strings.Split(goalStr[1], " ")
		numbers := make([]int, len(numberStr))
		goal, _ := strconv.Atoi(goalStr[0])
		for i, number := range numberStr {
			numbers[i], _ = strconv.Atoi(number)
		}

		newArr := make([]int, len(numbers)-1)
		copy(newArr, numbers[1:])
		task1 := Task1(goal, numbers[0], newArr, false)
		if task1 {
			task1sum = task1sum + goal
		}
		task2 := Task1(goal, numbers[0], newArr, true)
		if task2 {
			task2sum = task2sum + goal
		}
	}
	log.Println(task1sum)
	log.Println(task2sum)
}

func Task1(target, current int, arr []int, allowConcat bool) bool {
	if len(arr) == 0 {
		//finished
		return target == current
	}
	newArr := make([]int, len(arr)-1)
	copy(newArr, arr[1:])
	return Task1(target, current*arr[0], newArr, allowConcat) || Task1(target, current+arr[0], newArr, allowConcat) || (allowConcat && Task1(target, concatNums(current, arr[0]), newArr, allowConcat))
}

func concatNums(num1, num2 int) int {
	str1 := strconv.Itoa(num1)
	str2 := strconv.Itoa(num2)
	newNumStr := fmt.Sprintf("%s%s", str1, str2)
	newNum, _ := strconv.Atoi(newNumStr)
	return newNum
}

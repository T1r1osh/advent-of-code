package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadFile(path string) (list1 []int, list2 []int, err error) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return list1, list2, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Skipping line with invalid numbers:", line)
			continue
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return list1, list2, err
	}
	return list1, list2, nil
}

func CalculateDistance(list1 []int, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)
	distance := 0
	for index, item := range list1 {
		distance = distance + Abs(list2[index]-item)
	}
	return distance
}

func CalculateSimilarity(list1 []int, list2 []int) int {
	similarity := 0
	for _, item := range list1 {
		count := 0
		for _, list2item := range list2 {
			if list2item == item {
				count++
			}
		}
		similarity = similarity + (item * count)
	}
	return similarity
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	list1, list2, err := ReadFile("solutions/2024/Day01/input.txt")
	if err != nil {
		log.Fatal("Error:", err)
	}

	println("The distance is: ", CalculateDistance(list1, list2))

	println("The similarity score is: ", CalculateSimilarity(list1, list2))
}

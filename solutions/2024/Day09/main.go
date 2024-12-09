package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {

	file, _ := os.Open("solutions/2024/Day09/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	puzzleInput := scanner.Text()

	individualBlocks := []string{}

	for i, char := range puzzleInput {

		size, _ := strconv.Atoi(string(char))

		newArr := make([]string, size)
		character := ""
		if i%2 == 1 {
			character = "."
		} else {
			character = strconv.Itoa(i / 2)
		}
		for i := 0; i < size; i++ {
			newArr[i] = character
		}
		individualBlocks = append(individualBlocks, newArr...)
	}
	log.Println(Task1(individualBlocks))
	log.Println(Task2(individualBlocks))
}

func Task1(input []string) int {

	copyInput := make([]string, len(input))
	copy(copyInput, input)
	j := len(input) - 1
	for i := 0; i < j; i++ {
		if copyInput[i] == "." {
			copyInput[i], copyInput[j] = copyInput[j], copyInput[i]
			for {
				if copyInput[j] != "." {
					break
				}
				j--
			}
		}
	}

	return GetCheckSum(copyInput)
}

func Task2(input []string) int {
	copyInput := make([]string, len(input))
	copy(copyInput, input)
	fileMap := GetFileMap(copyInput)
	for i := len(fileMap) - 1; i >= 0; i-- {
		actualFile := strconv.Itoa(i)
		fileIndex := slices.Index(copyInput, actualFile)
		fileSize := fileMap[actualFile]
		spaceIndex := GetFirstIndexWithEnoughSpace(copyInput, fileSize)
		if spaceIndex != -1 && fileIndex > spaceIndex {
			for j := 0; j < fileSize; j++ {
				copyInput[spaceIndex+j], copyInput[fileIndex+j] = copyInput[fileIndex+j], copyInput[spaceIndex+j]
			}
		}
	}
	return GetCheckSum(copyInput)
}

func GetFirstIndexWithEnoughSpace(input []string, space int) int {
	lastChar := ""
	size := 0
	startindex := 0
	for i, in := range input {
		if in != lastChar {
			lastChar = in
			size = 0
			startindex = i
		}
		if in != "." {
			lastChar = ""
		}
		size++
		if size == space && lastChar == "." {
			return startindex
		}
	}
	return -1
}

func GetCheckSum(input []string) int {
	sum := 0
	for i, in := range input {
		val, _ := strconv.Atoi(in)
		sum = sum + i*val
	}
	return sum
}

// Describe a map between file id and size
func GetFileMap(input []string) map[string]int {
	lastChar := ""
	size := 0
	fileMap := map[string]int{}
	for _, in := range input {
		if in != lastChar {
			fileMap[lastChar] = size
			lastChar = in
			size = 0
		}
		size++
	}
	fileMap[lastChar] = size
	delete(fileMap, "")
	delete(fileMap, ".")
	return fileMap
}

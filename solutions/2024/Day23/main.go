package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	filePath := "solutions/2024/Day23/input.txt"
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	connections := map[string][]string{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")

		if _, ok := connections[line[0]]; ok {
			connections[line[0]] = append(connections[line[0]], line[1])
		} else {
			connections[line[0]] = make([]string, 0)
			connections[line[0]] = append(connections[line[0]], line[1])
		}
		if _, ok := connections[line[1]]; ok {
			connections[line[1]] = append(connections[line[1]], line[0])
		} else {
			connections[line[1]] = make([]string, 0)
			connections[line[1]] = append(connections[line[1]], line[0])
		}
	}
	log.Println(Task1(connections))

	cache := map[string]int{}
	for computer, connection := range connections {

		for i := 0; i < len(connection); i++ {
			cache[computer]++
		}
	}

	input := ReadFileLineByLine(filePath)
	lanMap := getLANMap(input)
	log.Println(findMaxCliques(lanMap))
}

func Task1(connections map[string][]string) int {
	resultMap := make(map[[3]string]struct{})
	for computer, connection := range connections {

		if computer[:1] != "t" {
			continue
		}

		for i := 0; i < len(connection); i++ {
			for j := i + 1; j < len(connection); j++ {
				// all combination
				cache := map[string]int{}
				cache[computer] = 1
				cache[connection[i]] = 1
				cache[connection[j]] = 1

				fasz := [2]string{connection[i], connection[j]}

				for _, connect := range fasz {
					cache[connect]++ // valszeg nem kell
					for _, val := range connections[connect] {
						if _, ok := cache[val]; ok {
							cache[val]++
						}
					}
				}
				//log.Println(cache)
				goodConn := true
				for _, nr := range cache {
					if nr != 3 {
						goodConn = false
					}
				}
				if goodConn {
					combination := [3]string{computer, connection[i], connection[j]}
					sort.Strings(combination[:])
					resultMap[combination] = struct{}{}
				}
			}
		}
	}
	//log.Println(resultMap)
	return len(resultMap)
}

func ReadFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

func getLANMap(input []string) map[string]map[string]struct{} {
	lanMap := make(map[string]map[string]struct{})

	for _, line := range input {
		comp1, comp2 := line[:2], line[3:]

		if _, ok := lanMap[comp1]; !ok {
			lanMap[comp1] = make(map[string]struct{})
		}

		if _, ok := lanMap[comp2]; !ok {
			lanMap[comp2] = make(map[string]struct{})
		}

		lanMap[comp1][comp2] = struct{}{}
		lanMap[comp2][comp1] = struct{}{}
	}
	return lanMap
}

func BronKerbosch(currentClique []string, yetToConsider []string, alreadyConsidered []string, lanMap map[string]map[string]struct{}, cliques [][]string) [][]string {
	if len(yetToConsider) == 0 && len(alreadyConsidered) == 0 {
		cliques = append(cliques, append([]string{}, currentClique...))
		return cliques
	}

	for index := 0; index < len(yetToConsider); {
		node := yetToConsider[index]
		newYetToConsider := []string{}
		newAlreadyConsidered := []string{}

		for _, n := range yetToConsider {
			if _, ok := lanMap[node][n]; ok {
				newYetToConsider = append(newYetToConsider, n)
			}
		}

		for _, n := range alreadyConsidered {
			if _, ok := lanMap[node][n]; ok {
				newAlreadyConsidered = append(newAlreadyConsidered, n)
			}
		}

		cliques = BronKerbosch(append(currentClique, node), newYetToConsider, newAlreadyConsidered, lanMap, cliques)

		yetToConsider = append(yetToConsider[:index], yetToConsider[index+1:]...)
		alreadyConsidered = append(alreadyConsidered, node)
	}
	return cliques
}

func findMaxCliques(lanMap map[string]map[string]struct{}) string {
	maxClique := []string{}
	allComputers := []string{}
	for key := range lanMap {
		allComputers = append(allComputers, key)
	}
	cliques := BronKerbosch([]string{}, allComputers, []string{}, lanMap, [][]string{})
	for _, c := range cliques {
		if len(c) > len(maxClique) {
			maxClique = c
		}
	}
	sort.Strings(maxClique)
	return strings.Join(maxClique, ",")
}

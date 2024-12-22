package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Secret int

func main() {

	file, _ := os.Open("solutions/2024/Day22/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var buyers []Secret
	for scanner.Scan() {
		line := scanner.Text()
		secretcode, _ := strconv.Atoi(line)
		buyers = append(buyers, Secret(secretcode))
	}

	log.Println(Task1(buyers))
	log.Println(Task2(buyers))

}

func Task1(buyers []Secret) int {
	sum := 0
	copyBuyers := make([]Secret, len(buyers))
	copy(copyBuyers, buyers)
	for _, buyer := range copyBuyers {
		for i := 0; i < 2000; i++ {
			buyer.Evolve()
		}
		sum += int(buyer)
	}
	return sum
}

func Task2(buyers []Secret) int {
	seqMap := map[[4]int]int{}
	copyBuyers := make([]Secret, len(buyers))
	copy(copyBuyers, buyers)
	for _, buyer := range copyBuyers {
		changes := []int{}
		init := (int(buyer)) % 10
		bananas := []int{init}
		contributed := map[[4]int]bool{}

		for i := 0; i < 2000; i++ {
			buyer.Evolve()
			lastDigit := (int(buyer)) % 10
			changes = append(changes, lastDigit-init)
			bananas = append(bananas, lastDigit)
			init = lastDigit
		}

		for i := 0; i < len(changes)-4; i++ {
			var key [4]int
			copy(key[:], changes[i:i+4])

			if contributed[key] {
				continue
			}

			contributed[key] = true
			seqMap[key] += bananas[i+4]
		}
	}
	var maxKey [4]int
	maxValue := 0

	for key, value := range seqMap {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	log.Println(maxKey, maxValue)
	return maxValue
}

func (s *Secret) Evolve() {
	value := *s * 64
	s.mix(value)
	s.prune()

	value = *s / 32
	s.mix(value)
	s.prune()

	value = *s * 2048
	s.mix(value)
	s.prune()
}

func (s *Secret) mix(value Secret) {
	*s = *s ^ Secret(value)
}

func (s *Secret) prune() {
	*s = *s % 16777216
}

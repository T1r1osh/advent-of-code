package main

import (
	"testing"
)

func TestDistance(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	distance := CalculateDistance(list1, list2)
	if distance != 11 {
		t.Error("Distance should be 11 insted of ", distance)
	}
}

func TestSimilarity(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	similarity := CalculateSimilarity(list1, list2)
	if similarity != 31 {
		t.Error("Similarity should be 31 insted of ", similarity)
	}

}

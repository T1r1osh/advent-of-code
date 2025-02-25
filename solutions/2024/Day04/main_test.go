package main

import "testing"

func TestTask1(t *testing.T) {

	grid := [][]string{
		{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
		{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
		{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
		{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
		{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
		{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
		{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
		{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
		{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
	}

	grid2 := [][]string{
		{".", "M", ".", "S", ".", ".", ".", ".", ".", "."},
		{".", ".", "A", ".", ".", "M", "S", "M", "S", "."},
		{".", "M", ".", "S", ".", "M", "A", "A", ".", "."},
		{".", ".", "A", ".", "A", "S", "M", "S", "M", "."},
		{".", "M", ".", "S", ".", "M", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"S", ".", "S", ".", "S", ".", "S", ".", "S", "."},
		{".", "A", ".", "A", ".", "A", ".", "A", ".", "."},
		{"M", ".", "M", ".", "M", ".", "M", ".", "M", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	task1 := Task1(grid)
	if task1 != 18 {
		t.Error("Result should be 18 instead of ", task1)
	}

	task2 := Task2(grid2)
	if task2 != 9 {
		t.Error("Result should be 9 instead of ", task2)
	}

}

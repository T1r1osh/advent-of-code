package main

import "testing"

func TestSafe(t *testing.T) {

	safeinput := []int{7, 6, 4, 2, 1}
	unsafeinput := []int{1, 2, 7, 8, 9}
	unsafeinput2 := []int{9, 7, 6, 2, 1}
	unsafeinput3 := []int{1, 3, 2, 4, 5}
	unsafeinput4 := []int{8, 6, 4, 4, 1}
	safeinput2 := []int{1, 3, 6, 7, 9}

	if !isSafeFloor(safeinput) {
		t.Error("This floor should be safe", safeinput)
	}

	if isSafeFloor(unsafeinput) {
		t.Error("This floor should be safe", unsafeinput)
	}

	if isSafeFloor(unsafeinput2) {
		t.Error("This floor should be safe", unsafeinput2)
	}

	if isSafeFloor(unsafeinput3) {
		t.Error("This floor should be safe", unsafeinput3)
	}

	if isSafeFloor(unsafeinput4) {
		t.Error("This floor should be safe", unsafeinput4)
	}

	if !isSafeFloor(safeinput2) {
		t.Error("This floor should be safe", safeinput2)
	}

	if !isSafeFloorWithDampener(unsafeinput4) {
		t.Error("This floor should be safe", unsafeinput4)
	}

	if !isSafeFloorWithDampener(unsafeinput3) {
		t.Error("This floor should be safe", unsafeinput3)
	}

}

package main

import "testing"

func TestSafe(t *testing.T) {
	answer := task1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	if answer != 161 {
		t.Error("The result should be 161 instead of ", answer)
	}

}

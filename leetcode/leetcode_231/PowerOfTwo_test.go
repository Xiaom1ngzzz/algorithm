package leetcode231

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	num1 := 32
	num2 := 33

	res1 := isPowerOfTwo(num1)
	res2 := isPowerOfTwo(num2)

	if res1 != true {
		t.Errorf("unexpected: %v , got: %v", true, res1)
	}

	if res2 != false {
		t.Errorf("unexpected: %v , got: %v", false, res2)
	}
}

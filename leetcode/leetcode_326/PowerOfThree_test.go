package leetcode326

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	nums := []int{3, 9, 27, 54, 81, 91, 127}
	resB := []bool{true, true, true, false, true, false, false}

	for i, v := range nums {
		res := isPowerOfThree(v)
		if res != resB[i] {
			t.Errorf("nums is %d, want: %v, got: %v", v, resB[i], res)
		}
	}
}

package edfe05a1d45c4ea89101d936cac32469

import "testing"

func TestSmallSum(t *testing.T) {
	arr := []int{1, 3, 5, 2, 4, 6}

	if smallSum(arr, 0, len(arr)-1) != 27 {
		t.Fatal("出错了")
	}
}

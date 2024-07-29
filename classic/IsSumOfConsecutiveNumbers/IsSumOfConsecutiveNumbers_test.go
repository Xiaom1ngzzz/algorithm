package issumofconsecutivenumbers

import "testing"

func TestIsSumOfConsecutiveNumbers(t *testing.T) {
	for num := 1; num < 200; num++ {
		res1 := is1(num)
		res2 := is2(num)

		if res1 != res2 {
			t.Errorf("%d 出错了", num)
		}
	}
}

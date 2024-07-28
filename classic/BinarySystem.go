package classic

import (
	"fmt"
	"unsafe"
)

func printBinary(num int) {
	// 下面这句写法，可以改成 :
	// (a & (1 << i)) != 0 ? "1" : "0"
	// 但不可以改成 :
	// (a & (1 << i)) == 1 ? "1" : "0"
	// 因为a如果第i位有1，那么(a & (1 << i))是2的i次方，而不一定是1
	// 比如，a = 0010011
	// a的第0位是1，第1位是1，第4位是1
	// (a & (1<<4)) == 16（不是1），说明a的第4位是1状态
	bitSize := int(unsafe.Sizeof(num)) * 8
	fmt.Printf("bitSize = %d\n", bitSize)
	for i := bitSize; i >= 0; i-- {
		if num&(1<<i) == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
	}
	fmt.Println()
}

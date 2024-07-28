package leetcode912

// i位置的数，向上调整大根堆
func heapInsert(arr []int, i int) {
	for arr[i] > arr[(i-1)/2] { // 满足两个条件：1. 当前节点有父节点；2. 当前节点的值大于父节点的值；
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i] // 将当前节点与其父节点进行交换
		i = (i - 1) / 2                             // 当前节点的位置变为其父节点位置
	}
}

// i位置的数，向下调整大根堆
func heapify(arr []int, i int, size int) {
	l := i*2 + 1   // 左孩子的位置
	for l < size { // 做孩子的位置应该小于堆的大小
		// 左孩子 l
		// 右孩子 l+1
		best := 0
		if l+1 < size && arr[l+1] > arr[l] {
			best = l + 1 // 右孩子
		} else {
			best = l // 左孩子
		}
		if arr[best] <= arr[i] {
			best = i
		}

		if best == i {
			break
		}
		swap(arr, best, i)
		i = best
		l = i*2 + 1
	}
}

func heapSort1(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		heapInsert(arr, i)
	}
	// fmt.Println("大根堆：", arr)

	size := n
	for size > 1 {
		size--
		swap(arr, 0, size)
		heapify(arr, 0, size)
	}
}

func heapSort2(arr []int) {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
	// fmt.Println("大根堆：", arr)

	size := n
	for size > 1 {
		size--
		swap(arr, 0, size)
		heapify(arr, 0, size)
	}
}

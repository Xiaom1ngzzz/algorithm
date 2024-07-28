package leetcode912

// swap 数组中交换i和j位置的数
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 选择排序
func selectionSort(arr []int) {
	minIndex := 0
	for i := 0; i < len(arr); i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

// 冒泡排序
func bubbleSort(arr []int) {
	for end := len(arr) - 1; end > 0; end-- {
		for j := 0; j < end; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 插入排序
func insertionSort(arr []int) {
	// 0 ~ 0
	// 0 ~ 1
	// ...
	// 0 ~ n-1
	for i := 1; i < len(arr); i++ {
		// 0 ~ i-1已经有序了，新的数arr[i]，向左看
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}

// func randomArray(n, v int) []int {
// 	arr := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		arr = append(arr, int(rand.Float64()*float64(v))+1)
// 	}
// 	return arr
// }

func copyArray(arr []int) []int {
	ans := make([]int, len(arr))
	copy(ans, arr)
	// fmt.Printf("%p %p\n", arr, ans)
	return ans
}

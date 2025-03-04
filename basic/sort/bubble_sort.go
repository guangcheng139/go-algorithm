package sort

// BubbleSort 冒泡排序
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// 稳定性: 稳定
func BubbleSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	for i := 0; i < n-1; i++ {
		// 提前退出标志
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 发生了交换，标记为 true
				swapped = true
			}
		}

		// 如果没有发生交换，表示数组已经有序，可以提前退出
		if !swapped {
			break
		}
	}

	return arr
}

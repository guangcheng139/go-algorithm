package sort

// QuickSort 快速排序
// 时间复杂度: 平均 O(nlogn)，最坏 O(n^2)
// 空间复杂度: O(logn) ~ O(n)
// 稳定性: 不稳定
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	quickSortRecursive(arr, 0, len(arr)-1)
	return arr
}

func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		// 获取分区点
		pivot := partition(arr, low, high)

		// 分治递归
		quickSortRecursive(arr, low, pivot-1)  // 左半部分
		quickSortRecursive(arr, pivot+1, high) // 右半部分
	}
}

// partition 分区函数
func partition(arr []int, low, high int) int {
	// 选择最右边的元素作为基准点
	pivot := arr[high]

	// i 是小于基准点元素的边界
	i := low - 1

	// 遍历数组，将小于基准点的元素放到左边
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			// 交换 arr[i] 和 arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 将基准点放到正确的位置
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// 返回基准点的索引
	return i + 1
}

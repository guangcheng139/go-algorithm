package sorting

// QuickSort 对整数切片进行快速排序
func QuickSort(arr []int) []int {
	// 创建一个副本，避免修改原始切片
	result := make([]int, len(arr))
	copy(result, arr)

	// 调用递归快速排序函数
	quickSortRecursive(result, 0, len(result)-1)

	return result
}

// quickSortRecursive 递归实现快速排序
func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		// 获取分区点
		pivotIndex := partition(arr, low, high)

		// 递归排序分区点左侧的元素
		quickSortRecursive(arr, low, pivotIndex-1)

		// 递归排序分区点右侧的元素
		quickSortRecursive(arr, pivotIndex+1, high)
	}
}

// partition 将数组分区，并返回分区点索引
func partition(arr []int, low, high int) int {
	// 选择最右边的元素作为基准值
	pivot := arr[high]

	// i 是小于基准值的元素的最后位置
	i := low - 1

	// 遍历从low到high-1的元素
	for j := low; j < high; j++ {
		// 如果当前元素小于或等于基准值
		if arr[j] <= pivot {
			// 增加i并交换元素
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 将基准值放到正确的位置
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// 返回基准值的索引
	return i + 1
}

// QuickSortWithRandomPivot 使用随机选择基准值的快速排序
func QuickSortWithRandomPivot(arr []int) []int {
	// 创建一个副本，避免修改原始切片
	result := make([]int, len(arr))
	copy(result, arr)

	// 调用递归快速排序函数
	quickSortWithRandomPivotRecursive(result, 0, len(result)-1)

	return result
}

// quickSortWithRandomPivotRecursive 使用随机基准值的递归快速排序
func quickSortWithRandomPivotRecursive(arr []int, low, high int) {
	if low < high {
		// 获取分区点
		pivotIndex := randomPartition(arr, low, high)

		// 递归排序分区点左侧的元素
		quickSortWithRandomPivotRecursive(arr, low, pivotIndex-1)

		// 递归排序分区点右侧的元素
		quickSortWithRandomPivotRecursive(arr, pivotIndex+1, high)
	}
}

// randomPartition 随机选择基准值并分区
func randomPartition(arr []int, low, high int) int {
	// 随机选择一个索引作为基准值
	randomIndex := low + (high-low)/2 // 这里使用中间元素，也可以使用随机数生成器

	// 将随机选择的元素与最后一个元素交换
	arr[randomIndex], arr[high] = arr[high], arr[randomIndex]

	// 调用标准分区函数
	return partition(arr, low, high)
}

// QuickSortInPlace 原地快速排序（不创建新切片）
func QuickSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}

	quickSortRecursive(arr, 0, len(arr)-1)
}

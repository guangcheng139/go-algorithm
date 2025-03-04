package sort

// MergeSort 归并排序
// 时间复杂度: O(nlogn)
// 空间复杂度: O(n)
// 稳定性: 稳定
func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	// 创建结果数组
	result := make([]int, n)

	// 递归排序
	mergeSortRecursive(arr, result, 0, n-1)

	return arr
}

func mergeSortRecursive(arr, result []int, start, end int) {
	// 递归终止条件
	if start >= end {
		return
	}

	// 计算中间点
	mid := start + (end-start)/2

	// 递归排序左右两部分
	mergeSortRecursive(arr, result, start, mid)
	mergeSortRecursive(arr, result, mid+1, end)

	// 合并两个有序数组
	merge(arr, result, start, mid, end)
}

// merge 合并两个有序数组
func merge(arr, result []int, start, mid, end int) {
	// 左半部分起始索引
	i := start
	// 右半部分起始索引
	j := mid + 1
	// 临时数组索引
	k := start

	// 比较左右两部分的元素，将较小的放入结果数组
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			result[k] = arr[i]
			i++
		} else {
			result[k] = arr[j]
			j++
		}
		k++
	}

	// 处理左半部分剩余元素
	for i <= mid {
		result[k] = arr[i]
		i++
		k++
	}

	// 处理右半部分剩余元素
	for j <= end {
		result[k] = arr[j]
		j++
		k++
	}

	// 将排序结果复制回原数组
	for i := start; i <= end; i++ {
		arr[i] = result[i]
	}
}

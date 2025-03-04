package search

// BinarySearch 二分查找（迭代实现）
// 在有序数组中查找目标值，返回目标值的索引，如果不存在则返回-1
// 时间复杂度: O(logn)
// 空间复杂度: O(1)
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // 找到目标值
		} else if arr[mid] < target {
			left = mid + 1 // 目标值在右半部分
		} else {
			right = mid - 1 // 目标值在左半部分
		}
	}

	return -1 // 未找到目标值
}

// BinarySearchRecursive 二分查找（递归实现）
// 在有序数组中查找目标值，返回目标值的索引，如果不存在则返回-1
// 时间复杂度: O(logn)
// 空间复杂度: O(logn)
func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchHelper(arr, target, 0, len(arr)-1)
}

func binarySearchHelper(arr []int, target, left, right int) int {
	// 基本情况：未找到目标值
	if left > right {
		return -1
	}

	// 计算中间索引
	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid // 找到目标值
	} else if arr[mid] < target {
		// 目标值在右半部分
		return binarySearchHelper(arr, target, mid+1, right)
	} else {
		// 目标值在左半部分
		return binarySearchHelper(arr, target, left, mid-1)
	}
}

// FirstOccurrence 查找第一个等于给定值的元素
// 时间复杂度: O(logn)
// 空间复杂度: O(1)
func FirstOccurrence(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid    // 记录当前找到的位置
			right = mid - 1 // 继续在左边查找
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// LastOccurrence 查找最后一个等于给定值的元素
// 时间复杂度: O(logn)
// 空间复杂度: O(1)
func LastOccurrence(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid   // 记录当前找到的位置
			left = mid + 1 // 继续在右边查找
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

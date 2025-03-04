package searching

// BinarySearchIterative 使用迭代方式实现二分查找
// 返回目标值在数组中的索引，如果未找到则返回-1
// 注意：数组必须已排序
func BinarySearchIterative(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		// 计算中间索引，避免整数溢出
		mid := left + (right-left)/2

		// 如果找到目标值
		if arr[mid] == target {
			return mid
		}

		// 如果目标值在左半部分
		if arr[mid] > target {
			right = mid - 1
		} else { // 如果目标值在右半部分
			left = mid + 1
		}
	}

	// 未找到目标值
	return -1
}

// BinarySearchRecursive 使用递归方式实现二分查找
// 返回目标值在数组中的索引，如果未找到则返回-1
// 注意：数组必须已排序
func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchRecursiveHelper(arr, target, 0, len(arr)-1)
}

// binarySearchRecursiveHelper 递归二分查找的辅助函数
func binarySearchRecursiveHelper(arr []int, target, left, right int) int {
	// 基本情况：未找到目标值
	if left > right {
		return -1
	}

	// 计算中间索引，避免整数溢出
	mid := left + (right-left)/2

	// 如果找到目标值
	if arr[mid] == target {
		return mid
	}

	// 如果目标值在左半部分
	if arr[mid] > target {
		return binarySearchRecursiveHelper(arr, target, left, mid-1)
	}

	// 如果目标值在右半部分
	return binarySearchRecursiveHelper(arr, target, mid+1, right)
}

// FindFirstOccurrence 查找目标值在排序数组中第一次出现的位置
// 如果未找到则返回-1
func FindFirstOccurrence(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		// 如果找到目标值
		if arr[mid] == target {
			result = mid    // 记录当前位置
			right = mid - 1 // 继续在左侧查找
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

// FindLastOccurrence 查找目标值在排序数组中最后一次出现的位置
// 如果未找到则返回-1
func FindLastOccurrence(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		// 如果找到目标值
		if arr[mid] == target {
			result = mid   // 记录当前位置
			left = mid + 1 // 继续在右侧查找
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

// SearchInRotatedSortedArray 在旋转排序数组中查找目标值
// 返回目标值在数组中的索引，如果未找到则返回-1
func SearchInRotatedSortedArray(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		// 如果找到目标值
		if arr[mid] == target {
			return mid
		}

		// 检查哪一部分是有序的
		if arr[left] <= arr[mid] { // 左半部分有序
			// 检查目标值是否在左半部分
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1 // 在左半部分查找
			} else {
				left = mid + 1 // 在右半部分查找
			}
		} else { // 右半部分有序
			// 检查目标值是否在右半部分
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1 // 在右半部分查找
			} else {
				right = mid - 1 // 在左半部分查找
			}
		}
	}

	// 未找到目标值
	return -1
}

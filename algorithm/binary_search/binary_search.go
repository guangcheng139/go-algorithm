package binary_search

// SearchInRotatedSortedArray 搜索旋转排序数组
// 整数数组 nums 按升序排列，数组中的值互不相同。在传递给函数之前，nums 在预先未知的某个下标 k 上进行了旋转，
// 使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]。
// 例如, [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2]。
// 搜索 nums 中的 target，如果存在则返回下标，否则返回 -1。
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid // 找到目标值
		}

		// 判断哪一部分是有序的
		if nums[left] <= nums[mid] { // 左半部分有序
			// 判断目标值是否在左半部分的有序区间内
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // 目标在左半部分
			} else {
				left = mid + 1 // 目标在右半部分
			}
		} else { // 右半部分有序
			// 判断目标值是否在右半部分的有序区间内
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // 目标在右半部分
			} else {
				right = mid - 1 // 目标在左半部分
			}
		}
	}

	return -1 // 未找到目标值
}

// FindMinimumInRotatedSortedArray 寻找旋转排序数组中的最小值
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2]。
// 请找出其中最小的元素。
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func FindMinimumInRotatedSortedArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	// 如果数组没有旋转，直接返回第一个元素
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left < right {
		mid := left + (right-left)/2

		// 如果中间元素大于右边元素，说明最小值在右半部分
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

// SearchRange 在排序数组中查找元素的第一个和最后一个位置
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值，返回 [-1, -1]。
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	if len(nums) == 0 {
		return result
	}

	// 查找第一个位置
	result[0] = findFirstPosition(nums, target)

	// 如果第一个位置不存在，则最后一个位置也不存在
	if result[0] == -1 {
		return result
	}

	// 查找最后一个位置
	result[1] = findLastPosition(nums, target)

	return result
}

// findFirstPosition 查找目标值的第一个位置
func findFirstPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid
			right = mid - 1 // 继续向左查找
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// findLastPosition 查找目标值的最后一个位置
func findLastPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid
			left = mid + 1 // 继续向右查找
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// PeakIndexInMountainArray 山脉数组的峰顶索引
// 符合下列属性的数组 arr 称为 山脉数组 ：
// arr.length >= 3
// 存在 i（0 < i < arr.length - 1）使得：
//
//	arr[0] < arr[1] < ... arr[i-1] < arr[i]
//	arr[i] > arr[i+1] > ... > arr[arr.length - 1]
//
// 给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func PeakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2

		// 如果中间元素小于右边元素，说明峰顶在右边
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// FindPeakElement 寻找峰值
// 峰值元素是指其值大于左右相邻值的元素。
// 给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。
// 数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
// 你可以假设 nums[-1] = nums[n] = -∞。
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// 如果中间元素小于右边元素，说明峰值在右边
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

package array

import (
	"sort"
)

// 两数之和：给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 时间复杂度：O(n)，空间复杂度：O(n)
func TwoSum(nums []int, target int) []int {
	// 创建一个map，用于存储每个元素及其索引
	numMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算与当前元素配对的另一个元素
		complement := target - num

		// 如果这个配对元素在map中，则找到了答案
		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}

		// 将当前元素及其索引添加到map中
		numMap[num] = i
	}

	// 如果没有找到答案，返回空切片
	return []int{}
}

// 三数之和：给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
// 请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 时间复杂度：O(n^2)，空间复杂度：O(n)
func ThreeSum(nums []int) [][]int {
	result := [][]int{}

	// 处理边界情况
	n := len(nums)
	if n < 3 {
		return result
	}

	// 对数组进行排序
	sort.Ints(nums)

	// 固定第一个数字，然后使用双指针寻找另外两个数字
	for i := 0; i < n-2; i++ {
		// 跳过重复的元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 如果第一个数字大于0，后面的数字都大于0，三数之和不可能为0
		if nums[i] > 0 {
			break
		}

		// 使用双指针查找另外两个数字
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				// 如果和小于0，移动左指针
				left++
			} else if sum > 0 {
				// 如果和大于0，移动右指针
				right--
			} else {
				// 找到一个三元组
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 跳过重复的元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 继续寻找新的组合
				left++
				right--
			}
		}
	}

	return result
}

// 合并两个有序数组：给你两个按非递减顺序排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n，分别表示 nums1 和 nums2 中的元素数目。
// 请你合并 nums2 到 nums1 中，使合并后的数组同样按非递减顺序排列。
// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0，应忽略。nums2 的长度为 n。
// 时间复杂度：O(m+n)，空间复杂度：O(1)
func Merge(nums1 []int, m int, nums2 []int, n int) {
	// 从后往前合并，避免覆盖nums1中的元素
	i, j, k := m-1, n-1, m+n-1

	// 从后往前遍历两个数组
	for j >= 0 {
		// 如果nums1的当前元素大于nums2的当前元素，或者nums1已经用完
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// 旋转数组：给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
// 时间复杂度：O(n)，空间复杂度：O(1)
func Rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 || k%n == 0 {
		return
	}

	// 取模，以处理k大于n的情况
	k = k % n

	// 三次翻转
	// 1. 翻转整个数组
	reverse(nums, 0, n-1)
	// 2. 翻转前k个元素
	reverse(nums, 0, k-1)
	// 3. 翻转剩余元素
	reverse(nums, k, n-1)
}

// 辅助函数：翻转数组的一部分
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 移动零：给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 时间复杂度：O(n)，空间复杂度：O(1)
func MoveZeroes(nums []int) {
	// 使用双指针，一个指向要放置非零元素的位置，一个用于遍历数组
	nonZeroPos := 0

	// 遍历数组，将非零元素放到nonZeroPos位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroPos] = nums[i]
			nonZeroPos++
		}
	}

	// 将剩余位置填充为0
	for i := nonZeroPos; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 盛最多水的容器：给定一个长度为 n 的整数数组 height。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 时间复杂度：O(n)，空间复杂度：O(1)
func MaxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1

	for left < right {
		// 计算当前区域的面积
		width := right - left
		// 取两侧高度的较小值
		h := min(height[left], height[right])
		area := width * h

		// 更新最大面积
		maxArea = max(maxArea, area)

		// 移动指针
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// 最大子数组和：给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组 是数组中的一个连续部分。
// 时间复杂度：O(n)，空间复杂度：O(1)
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]     // 全局最大和
	currentSum := nums[0] // 当前子数组的和

	for i := 1; i < len(nums); i++ {
		// 如果当前子数组和为负数，那么对后面的元素没有增益，重新开始
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			// 否则，将当前元素加入到子数组中
			currentSum += nums[i]
		}

		// 更新最大和
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// 辅助函数：取两个数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 辅助函数：取两个数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

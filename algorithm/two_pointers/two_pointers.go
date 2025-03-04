package two_pointers

import "sort"

// 排序数组的两数之和：给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
// 函数应该返回这两个下标值 index1 和 index2，其中 index1 < index2。
// 时间复杂度：O(n)，空间复杂度：O(1)
func TwoSumSorted(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			// 题目要求返回的索引是1-indexed
			return []int{left + 1, right + 1}
		} else if sum < target {
			// 如果和小于目标值，增大左指针
			left++
		} else {
			// 如果和大于目标值，减小右指针
			right--
		}
	}

	// 没有找到答案
	return []int{-1, -1}
}

// 删除排序数组中的重复项：给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次。
// 返回移除后数组的新长度。
// 时间复杂度：O(n)，空间复杂度：O(1)
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用i记录不重复元素应该放置的位置
	i := 0

	// 从第二个元素开始遍历
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			// 找到一个新的不重复元素
			i++
			nums[i] = nums[j]
		}
	}

	// 返回新数组的长度
	return i + 1
}

// 接雨水：给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 时间复杂度：O(n)，空间复杂度：O(1)
func Trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	for left < right {
		// 更新左右两侧最大高度
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}

		// 计算当前位置能接的雨水
		if leftMax < rightMax {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}

	return result
}

// 颜色分类（荷兰国旗问题）：给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序。
// 使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 此题中，使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 时间复杂度：O(n)，空间复杂度：O(1)
func SortColors(nums []int) {
	// 定义三个指针
	// p0: 指向0序列的末尾
	// p2: 指向2序列的开头
	// curr: 当前考察的元素
	p0, curr, p2 := 0, 0, len(nums)-1

	for curr <= p2 {
		if nums[curr] == 0 {
			// 当前元素是0，交换到左侧
			nums[p0], nums[curr] = nums[curr], nums[p0]
			p0++
			curr++
		} else if nums[curr] == 2 {
			// 当前元素是2，交换到右侧
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
			// 注意这里不增加curr，因为交换来的元素需要重新检查
		} else {
			// 当前元素是1，不需要交换
			curr++
		}
	}
}

// 最接近的三数之和：给定一个包括 n 个整数的数组 nums 和一个目标值 target。
// 找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。
// 假定每组输入只存在唯一答案。
// 时间复杂度：O(n^2)，空间复杂度：O(1)
func ThreeSumClosest(nums []int, target int) int {
	// 先排序
	sort.Ints(nums)

	// 初始化答案为前三个元素的和
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 使用双指针
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// 更新最接近的和
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				// 如果和等于target，直接返回
				return sum
			}
		}
	}

	return closest
}

// 反转字符串：编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
// 时间复杂度：O(n)，空间复杂度：O(1)
func ReverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 判断回文字符串：给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// 说明：本题中，我们将空字符串定义为有效的回文串。
// 时间复杂度：O(n)，空间复杂度：O(1)
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// 跳过非字母数字字符
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		// 比较字符（忽略大小写）
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

// 辅助函数：计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 辅助函数：判断字符是否为字母或数字
func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// 辅助函数：将字符转换为小写
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}

package sliding_window

import (
	"math"
)

// MaxSlidingWindow 滑动窗口最大值
// 给定一个数组 nums 和一个滑动窗口大小 k，找出所有滑动窗口里的最大值
// 时间复杂度: O(n)
// 空间复杂度: O(k)
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}

	n := len(nums)
	result := make([]int, n-k+1)
	deque := make([]int, 0) // 双端队列，存储的是索引

	for i := 0; i < n; i++ {
		// 移除队列中所有小于当前元素的值（因为它们不可能是最大值）
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 添加当前索引到队列
		deque = append(deque, i)

		// 移除超出窗口范围的索引
		if deque[0] <= i-k {
			deque = deque[1:]
		}

		// 当形成第一个完整窗口时开始记录结果
		if i >= k-1 {
			result[i-k+1] = nums[deque[0]]
		}
	}

	return result
}

// MinSubArrayLen 长度最小的子数组
// 给定一个正整数数组 nums 和一个正整数 target，找出该数组中满足其和 ≥ target 的长度最小的连续子数组，并返回其长度
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	minLen := math.MaxInt32
	left, sum := 0, 0

	for right := 0; right < n; right++ {
		sum += nums[right]

		// 当窗口内元素和大于等于目标值时，尝试收缩窗口
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0 // 没有找到符合条件的子数组
	}
	return minLen
}

// CharacterReplacement 替换后的最长重复字符
// 给你一个字符串 s 和一个整数 k。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。
// 在执行最多 k 次更改后，返回包含相同字母的最长子串的长度。
// 时间复杂度: O(n)
// 空间复杂度: O(1)，因为字符集固定
func CharacterReplacement(s string, k int) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	count := make([]int, 26) // 记录窗口内每个字符的数量
	maxCount := 0            // 窗口内出现最多的字符的数量
	left := 0
	result := 0

	for right := 0; right < n; right++ {
		// 更新当前字符的数量
		count[s[right]-'A']++

		// 更新窗口内出现最多的字符的数量
		maxCount = max(maxCount, count[s[right]-'A'])

		// 当窗口内需要替换的字符数量 > k 时，收缩窗口
		// 窗口大小 - 出现最多字符的数量 = 需要被替换的字符数量
		if right-left+1-maxCount > k {
			count[s[left]-'A']--
			left++
		}

		// 更新结果
		result = max(result, right-left+1)
	}

	return result
}

// FindAnagrams 找到字符串中所有字母异位词
// 给定两个字符串 s 和 p，找到 s 中所有 p 的字母异位词的子串，返回这些子串的起始索引
// 时间复杂度: O(n)
// 空间复杂度: O(1)，因为字符集固定
func FindAnagrams(s string, p string) []int {
	ns, np := len(s), len(p)
	if ns < np {
		return []int{}
	}

	pCount := make([]int, 26) // p中每个字符的数量
	sCount := make([]int, 26) // 窗口内每个字符的数量
	result := []int{}

	// 统计p中每个字符的数量
	for i := 0; i < np; i++ {
		pCount[p[i]-'a']++
	}

	// 初始窗口
	for i := 0; i < np; i++ {
		sCount[s[i]-'a']++
	}

	// 检查初始窗口是否是异位词
	if isEqual(pCount, sCount) {
		result = append(result, 0)
	}

	// 滑动窗口
	for i := np; i < ns; i++ {
		// 移除窗口最左侧字符
		sCount[s[i-np]-'a']--
		// 添加当前字符
		sCount[s[i]-'a']++

		// 检查当前窗口是否是异位词
		if isEqual(pCount, sCount) {
			result = append(result, i-np+1)
		}
	}

	return result
}

// isEqual 检查两个数组是否相等
func isEqual(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// min 返回两个数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max 返回两个数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

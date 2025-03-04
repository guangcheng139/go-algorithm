package string

import "strings"

// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。
// 如果不存在，则返回 -1。
// 时间复杂度：O(n*m)，其中n是haystack的长度，m是needle的长度
// 空间复杂度：O(1)
func StrStr(haystack string, needle string) int {
	// 如果needle为空，根据题意返回0
	if len(needle) == 0 {
		return 0
	}

	var i, j int
	// i只需要遍历到len(haystack)-len(needle)，后面的字符即使匹配也不够needle的长度
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		// 尝试从haystack的i位置开始匹配needle
		for j = 0; j < len(needle); j++ {
			// 如果当前字符不匹配，就退出内循环
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 如果j等于needle的长度，说明完全匹配到了needle
		if len(needle) == j {
			return i
		}
	}
	// 没有找到匹配，返回-1
	return -1
}

// 给定一个非空的字符串 s，检查它是否可以由它的一个子串重复多次构成。
// 例如，"abab"可以由"ab"重复两次构成，所以返回true
// 时间复杂度：O(n^2)，其中n是字符串s的长度
// 空间复杂度：O(n)，主要是strings.Repeat可能需要的空间
func RepeatedSubstringPattern(s string) bool {
	// 如果字符串长度小于2，不可能由重复子串构成
	if len(s) < 2 {
		return false
	}

	// 可能的重复子串长度不会超过字符串长度的一半
	for i := 0; i < len(s)/2; i++ {
		// 尝试以s[0:i+1]为重复单元
		str := s[0 : i+1]
		// 首先检查s的长度是否是str长度的整数倍
		if len(s)%len(str) == 0 {
			// 将str重复n次，看是否等于s
			if s == strings.Repeat(str, len(s)/len(str)) {
				return true
			}
		}
	}
	return false
}

// Sunday算法是一种改进的字符串匹配算法，比简单的暴力匹配更高效。
// 在haystack字符串中找出needle字符串出现的第一个位置。
// 时间复杂度：最坏情况O(n*m)，平均情况要好得多
// 空间复杂度：O(256)，使用了一个大小为256的数组记录字符位置
func StrStrSunday(haystack string, needle string) int {
	// 特殊情况处理：needle为空时，返回0
	if len(needle) == 0 {
		return 0
	}

	// 如果haystack的长度小于needle的长度，不可能匹配成功
	if len(haystack) < len(needle) {
		return -1
	}

	// 为每个字符计算在needle中最右出现的位置
	// 初始化为needle长度+1，表示字符不在needle中
	last := make([]int, 256)
	for i := 0; i < 256; i++ {
		last[i] = len(needle) + 1
	}

	// 更新needle中字符的最右位置
	for i := 0; i < len(needle); i++ {
		last[needle[i]] = len(needle) - i
	}

	// 开始匹配
	pos := 0 // 当前匹配位置
	for pos <= len(haystack)-len(needle) {
		// 尝试匹配
		i := 0
		for i < len(needle) && haystack[pos+i] == needle[i] {
			i++
		}

		// 如果完全匹配，返回位置
		if i == len(needle) {
			return pos
		}

		// 计算下一个匹配位置
		if pos+len(needle) >= len(haystack) {
			// 已经到达字符串末尾，无法继续匹配
			return -1
		}

		// 根据haystack[pos+len(needle)]计算跳转距离
		pos += last[haystack[pos+len(needle)]]
	}

	return -1
}

// KMP算法是一种高效的字符串匹配算法，避免了不必要的比较。
// 时间复杂度：O(n+m)，其中n是haystack的长度，m是needle的长度
// 空间复杂度：O(m)，需要一个next数组记录部分匹配表
func StrStrKMP(haystack string, needle string) int {
	// 特殊情况处理：needle为空时，返回0
	if len(needle) == 0 {
		return 0
	}

	// 如果haystack的长度小于needle的长度，不可能匹配成功
	if len(haystack) < len(needle) {
		return -1
	}

	// 构建KMP算法的部分匹配表（next数组）
	next := buildKMPNext(needle)

	i, j := 0, 0
	// i指向haystack，j指向needle
	for i < len(haystack) && j < len(needle) {
		// 如果j为-1或者当前字符匹配成功
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			// 匹配失败，j回退
			j = next[j]
		}

		// 如果完全匹配成功
		if j == len(needle) {
			return i - j
		}
	}

	return -1
}

// 构建KMP算法的部分匹配表
func buildKMPNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	i, j := 0, -1

	for i < len(pattern)-1 {
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}

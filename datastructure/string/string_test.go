package string

import (
	"testing"
)

// 测试StrStr函数
func TestStrStr(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"a", "", 0},
		{"mississippi", "issip", 4},
		{"aaa", "aaaa", -1},
		{"aaa", "aaa", 0},
	}

	// 运行测试用例
	for i, test := range tests {
		result := StrStr(test.haystack, test.needle)
		if result != test.expected {
			t.Errorf("测试%d失败: StrStr(%q, %q) = %d，期望值 %d",
				i, test.haystack, test.needle, result, test.expected)
		}
	}
}

// 测试RepeatedSubstringPattern函数
func TestRepeatedSubstringPattern(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		s        string
		expected bool
	}{
		{"abab", true},
		{"aba", false},
		{"abcabcabc", true},
		{"a", false},
		{"aa", true},
		{"abaababaab", true},
		{"abac", false},
		{"ababab", true},
	}

	// 运行测试用例
	for i, test := range tests {
		result := RepeatedSubstringPattern(test.s)
		if result != test.expected {
			t.Errorf("测试%d失败: RepeatedSubstringPattern(%q) = %v，期望值 %v",
				i, test.s, result, test.expected)
		}
	}
}

// 测试StrStrSunday函数
func TestStrStrSunday(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"a", "", 0},
		{"mississippi", "issip", 4},
		{"aaa", "aaaa", -1},
		{"aaa", "aaa", 0},
		{"babbbbbabb", "bbab", 5},
		{"GCATCGCAGAGAGTATACAGTACG", "GCAGAGAG", 5},
	}

	// 运行测试用例
	for i, test := range tests {
		result := StrStrSunday(test.haystack, test.needle)
		if result != test.expected {
			t.Errorf("测试%d失败: StrStrSunday(%q, %q) = %d，期望值 %d",
				i, test.haystack, test.needle, result, test.expected)
		}
	}
}

// 测试StrStrKMP函数
func TestStrStrKMP(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"a", "", 0},
		{"mississippi", "issip", 4},
		{"aaa", "aaaa", -1},
		{"aaa", "aaa", 0},
		{"ABABCABABA", "ABABC", 0},
		{"ABABDABACDABABCABAB", "ABABCABAB", 10},
	}

	// 运行测试用例
	for i, test := range tests {
		result := StrStrKMP(test.haystack, test.needle)
		if result != test.expected {
			t.Errorf("测试%d失败: StrStrKMP(%q, %q) = %d，期望值 %d",
				i, test.haystack, test.needle, result, test.expected)
		}
	}
}

// 性能比较：测试不同字符串匹配算法的执行时间
func BenchmarkStrStr(b *testing.B) {
	haystack := "mississippimississippimississippimississippimississippi"
	needle := "issip"
	for i := 0; i < b.N; i++ {
		StrStr(haystack, needle)
	}
}

func BenchmarkStrStrSunday(b *testing.B) {
	haystack := "mississippimississippimississippimississippimississippi"
	needle := "issip"
	for i := 0; i < b.N; i++ {
		StrStrSunday(haystack, needle)
	}
}

func BenchmarkStrStrKMP(b *testing.B) {
	haystack := "mississippimississippimississippimississippimississippi"
	needle := "issip"
	for i := 0; i < b.N; i++ {
		StrStrKMP(haystack, needle)
	}
}

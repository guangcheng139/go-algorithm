package greedy

import (
	"math"
	"sort"
	"testing"
)

func TestActivitySelection(t *testing.T) {
	tests := []struct {
		name     string
		start    []int
		end      []int
		expected int
	}{
		{
			name:     "空活动列表",
			start:    []int{},
			end:      []int{},
			expected: 0,
		},
		{
			name:     "单个活动",
			start:    []int{1},
			end:      []int{2},
			expected: 1,
		},
		{
			name:     "标准示例",
			start:    []int{1, 3, 0, 5, 8, 5},
			end:      []int{2, 4, 6, 7, 9, 9},
			expected: 4, // 可以选择的活动: [1,2], [3,4], [5,7], [8,9]
		},
		{
			name:     "所有活动冲突",
			start:    []int{1, 1, 1, 1},
			end:      []int{2, 2, 2, 2},
			expected: 1, // 只能选择一个活动
		},
		{
			name:     "没有冲突",
			start:    []int{1, 3, 5, 7},
			end:      []int{2, 4, 6, 8},
			expected: 4, // 可以选择所有活动
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ActivitySelection(tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("ActivitySelection() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHuffmanCoding(t *testing.T) {
	tests := []struct {
		name           string
		chars          []rune
		freq           []int
		expectedLen    map[rune]int // 仅测试编码长度
		expectedPrefix bool
	}{
		{
			name:           "空字符集",
			chars:          []rune{},
			freq:           []int{},
			expectedLen:    nil,
			expectedPrefix: true,
		},
		{
			name:           "单个字符",
			chars:          []rune{'a'},
			freq:           []int{100},
			expectedLen:    map[rune]int{'a': 0}, // 单个字符的编码长度为0
			expectedPrefix: true,
		},
		{
			name:           "标准示例",
			chars:          []rune{'a', 'b', 'c', 'd', 'e', 'f'},
			freq:           []int{5, 9, 12, 13, 16, 45},
			expectedLen:    map[rune]int{'a': 4, 'b': 4, 'c': 3, 'd': 3, 'e': 3, 'f': 1},
			expectedPrefix: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HuffmanCoding(tt.chars, tt.freq)

			// 对于空输入的特殊处理
			if len(tt.chars) == 0 {
				if result != nil {
					t.Errorf("HuffmanCoding() = %v, want nil", result)
				}
				return
			}

			// 检查编码长度
			for char, expectedLength := range tt.expectedLen {
				if code, ok := result[char]; ok {
					if len(code) != expectedLength {
						t.Errorf("HuffmanCoding() 字符 %c 的编码长度 = %v, want %v", char, len(code), expectedLength)
					}
				} else {
					t.Errorf("HuffmanCoding() 未找到字符 %c 的编码", char)
				}
			}

			// 验证前缀性质: 没有编码是另一个编码的前缀
			if tt.expectedPrefix {
				validatePrefixProperty(t, result)
			}
		})
	}
}

// 验证哈夫曼编码的前缀性质
func validatePrefixProperty(t *testing.T, codeMap map[rune]string) {
	codes := make([]string, 0, len(codeMap))
	for _, code := range codeMap {
		codes = append(codes, code)
	}

	// 按长度排序，短编码在前
	sort.Slice(codes, func(i, j int) bool {
		return len(codes[i]) < len(codes[j])
	})

	// 检查是否有编码是另一个编码的前缀
	for i := 0; i < len(codes); i++ {
		for j := i + 1; j < len(codes); j++ {
			if len(codes[i]) < len(codes[j]) && codes[j][:len(codes[i])] == codes[i] {
				t.Errorf("哈夫曼编码违反前缀性质: %s 是 %s 的前缀", codes[i], codes[j])
			}
		}
	}
}

func TestFractionalKnapsack(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		values   []int
		capacity int
		expected float64
	}{
		{
			name:     "空背包",
			weights:  []int{},
			values:   []int{},
			capacity: 10,
			expected: 0,
		},
		{
			name:     "零容量",
			weights:  []int{1, 2, 3},
			values:   []int{6, 10, 12},
			capacity: 0,
			expected: 0,
		},
		{
			name:     "标准示例",
			weights:  []int{10, 20, 30},
			values:   []int{60, 100, 120},
			capacity: 50,
			expected: 240.0, // 取整个项目1和2，以及项目3的一部分
		},
		{
			name:     "全部放入",
			weights:  []int{1, 2, 3},
			values:   []int{6, 10, 12},
			capacity: 10,
			expected: 28.0, // 可以放入所有物品
		},
		{
			name:     "部分放入",
			weights:  []int{5, 10, 15, 20},
			values:   []int{30, 20, 45, 80},
			capacity: 30,
			expected: 125.0, // 取整个项目1、3和项目4的一部分
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FractionalKnapsack(tt.weights, tt.values, tt.capacity)
			// 使用epsilon比较浮点数，避免浮点精度问题
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("FractionalKnapsack() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCoinChangeGreedy(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
		willWork bool // 表示贪心算法是否能得到最优解
	}{
		{
			name:     "空硬币列表",
			coins:    []int{},
			amount:   10,
			expected: -1,
			willWork: true,
		},
		{
			name:     "零金额",
			coins:    []int{1, 5, 10, 25},
			amount:   0,
			expected: 0,
			willWork: true,
		},
		{
			name:     "美国硬币系统",
			coins:    []int{1, 5, 10, 25},
			amount:   63,
			expected: 6, // 2个25分, 1个10分, 3个1分
			willWork: true,
		},
		{
			name:     "贪心不一定最优",
			coins:    []int{1, 3, 4},
			amount:   6,
			expected: 2, // 最优解是2个3分，而贪心会选择1个4分+2个1分
			willWork: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CoinChangeGreedy(tt.coins, tt.amount)

			// 仅验证贪心方法能够正确工作的情况
			if tt.willWork {
				if result != tt.expected {
					t.Errorf("CoinChangeGreedy() = %v, want %v", result, tt.expected)
				}
			} else {
				t.Logf("CoinChangeGreedy() = %v, 但这个测试用例贪心算法不一定得到最优解", result)
			}
		})
	}
}

func TestMinimumSpanningTree(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected int
	}{
		{
			name:     "空图",
			n:        0,
			edges:    [][]int{},
			expected: 0,
		},
		{
			name:     "单个节点",
			n:        1,
			edges:    [][]int{},
			expected: 0,
		},
		{
			name:     "标准示例",
			n:        4,
			edges:    [][]int{{0, 1, 10}, {0, 2, 6}, {0, 3, 5}, {1, 3, 15}, {2, 3, 4}},
			expected: 19, // 边 (0,3), (2,3), (0,1) 的权重和
		},
		{
			name:     "无法形成生成树",
			n:        5,
			edges:    [][]int{{0, 1, 1}, {1, 2, 2}, {3, 4, 3}},
			expected: -1, // 图不连通，无法形成生成树
		},
		{
			name:     "多余的边",
			n:        3,
			edges:    [][]int{{0, 1, 5}, {1, 2, 3}, {0, 2, 1}, {0, 1, 2}},
			expected: 3, // 边 (0,2), (0,1) 的权重和
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinimumSpanningTree(tt.n, tt.edges)
			if result != tt.expected {
				t.Errorf("MinimumSpanningTree() = %v, want %v", result, tt.expected)
			}
		})
	}
}

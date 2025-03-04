package greedy

import (
	"sort"
)

// ActivitySelection 活动选择问题
// 给定n个活动，每个活动有开始时间start[i]和结束时间end[i]。
// 你想参加尽可能多的活动，但前提是同一时间只能参加一个活动。
// 返回你能参加的最大活动数量。
// 时间复杂度: O(n log n)
// 空间复杂度: O(n)
func ActivitySelection(start []int, end []int) int {
	n := len(start)
	if n == 0 {
		return 0
	}

	// 创建活动结构体切片
	activities := make([]Activity, n)
	for i := 0; i < n; i++ {
		activities[i] = Activity{start[i], end[i]}
	}

	// 按结束时间排序
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].End < activities[j].End
	})

	count := 1 // 至少可以参加第一个活动
	lastEnd := activities[0].End

	// 贪心选择结束时间早的活动
	for i := 1; i < n; i++ {
		if activities[i].Start >= lastEnd {
			// 当前活动的开始时间晚于上一个选择的活动的结束时间
			count++
			lastEnd = activities[i].End
		}
	}

	return count
}

// Activity 活动结构体
type Activity struct {
	Start int
	End   int
}

// HuffmanCoding 哈夫曼编码
// 给定一组字符及其出现频率，构建哈夫曼树并返回编码结果。
// 时间复杂度: O(n log n)
// 空间复杂度: O(n)
func HuffmanCoding(chars []rune, freq []int) map[rune]string {
	n := len(chars)
	if n == 0 {
		return nil
	}

	// 创建叶子节点
	nodes := make([]*HuffmanNode, n)
	for i := 0; i < n; i++ {
		nodes[i] = &HuffmanNode{
			Char:  chars[i],
			Freq:  freq[i],
			Left:  nil,
			Right: nil,
		}
	}

	// 构建哈夫曼树
	root := buildHuffmanTree(nodes)

	// 生成编码
	codeMap := make(map[rune]string)
	generateCodes(root, "", codeMap)

	return codeMap
}

// HuffmanNode 哈夫曼树节点
type HuffmanNode struct {
	Char  rune
	Freq  int
	Left  *HuffmanNode
	Right *HuffmanNode
}

// buildHuffmanTree 构建哈夫曼树
func buildHuffmanTree(nodes []*HuffmanNode) *HuffmanNode {
	n := len(nodes)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return nodes[0]
	}

	// 复制原始节点数组，以便操作
	nodeList := make([]*HuffmanNode, n)
	copy(nodeList, nodes)

	for len(nodeList) > 1 {
		// 排序，找到频率最小的两个节点
		sort.Slice(nodeList, func(i, j int) bool {
			return nodeList[i].Freq < nodeList[j].Freq
		})

		// 取出频率最小的两个节点
		left := nodeList[0]
		right := nodeList[1]

		// 创建新的内部节点
		internal := &HuffmanNode{
			Char:  0, // 内部节点没有字符
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}

		// 移除已处理的两个节点，添加新的内部节点
		nodeList = append(nodeList[2:], internal)
	}

	// 返回根节点
	return nodeList[0]
}

// generateCodes 递归生成哈夫曼编码
func generateCodes(node *HuffmanNode, code string, codeMap map[rune]string) {
	if node == nil {
		return
	}

	// 叶子节点，保存编码
	if node.Left == nil && node.Right == nil {
		codeMap[node.Char] = code
		return
	}

	// 递归左子树，编码添加0
	generateCodes(node.Left, code+"0", codeMap)

	// 递归右子树，编码添加1
	generateCodes(node.Right, code+"1", codeMap)
}

// FractionalKnapsack 分数背包问题
// 给定n个物品，每个物品有重量weights[i]和价值values[i]。
// 背包的最大承重为capacity，求解放入背包的物品的最大价值。
// 与01背包不同，这里的物品可以部分放入背包。
// 时间复杂度: O(n log n)
// 空间复杂度: O(n)
func FractionalKnapsack(weights []int, values []int, capacity int) float64 {
	n := len(weights)
	if n == 0 || capacity == 0 {
		return 0
	}

	// 创建物品结构体切片
	items := make([]Item, n)
	for i := 0; i < n; i++ {
		items[i] = Item{
			Weight:         weights[i],
			Value:          values[i],
			ValuePerWeight: float64(values[i]) / float64(weights[i]),
		}
	}

	// 按单位价值排序（降序）
	sort.Slice(items, func(i, j int) bool {
		return items[i].ValuePerWeight > items[j].ValuePerWeight
	})

	totalValue := 0.0
	remainingCapacity := capacity

	// 贪心选择单位价值高的物品
	for i := 0; i < n; i++ {
		if items[i].Weight <= remainingCapacity {
			// 可以完全放入
			totalValue += float64(items[i].Value)
			remainingCapacity -= items[i].Weight
		} else {
			// 只能部分放入
			fraction := float64(remainingCapacity) / float64(items[i].Weight)
			totalValue += fraction * float64(items[i].Value)
			break // 背包已满
		}
	}

	return totalValue
}

// Item 物品结构体
type Item struct {
	Weight         int
	Value          int
	ValuePerWeight float64
}

// CoinChange 贪心算法解决零钱兑换问题（仅适用于某些特殊情况）
// 给定不同面额的硬币 coins 和一个总金额 amount，计算可以凑成总金额所需的最少的硬币个数。
// 注意：这种贪心算法只适用于特定的硬币系统，如美国硬币系统 {1, 5, 10, 25}
// 对于一般情况，应该使用动态规划（见dynamic_programming包）
// 时间复杂度: O(n log n + amount)
// 空间复杂度: O(n)
func CoinChangeGreedy(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 按面额从大到小排序
	sortedCoins := make([]int, len(coins))
	copy(sortedCoins, coins)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedCoins)))

	count := 0
	remaining := amount

	// 贪心选择最大面额的硬币
	for _, coin := range sortedCoins {
		// 使用当前面额的硬币数量
		count += remaining / coin
		// 剩余金额
		remaining %= coin

		if remaining == 0 {
			break
		}
	}

	// 如果无法凑成总金额，返回-1
	if remaining > 0 {
		return -1
	}

	return count
}

// MinimumSpanningTree Kruskal算法求最小生成树
// 给定一个无向带权图，求解最小生成树的总权重
// 输入：n为节点数量（节点编号从0到n-1），edges为边的列表，每条边包含[u, v, weight]
// 返回：最小生成树的总权重
// 时间复杂度: O(E log E)，其中E是边的数量
// 空间复杂度: O(V + E)，其中V是节点数量，E是边的数量
func MinimumSpanningTree(n int, edges [][]int) int {
	if n <= 1 || len(edges) == 0 {
		return 0
	}

	// 按边的权重排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// 初始化并查集
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	totalWeight := 0
	edgeCount := 0

	// Kruskal算法
	for _, edge := range edges {
		u, v, weight := edge[0], edge[1], edge[2]

		// 如果添加这条边不会形成环
		if find(parent, u) != find(parent, v) {
			// 合并两个集合
			union(parent, u, v)
			totalWeight += weight
			edgeCount++

			// 最小生成树有n-1条边
			if edgeCount == n-1 {
				break
			}
		}
	}

	// 如果无法形成最小生成树，返回-1
	if edgeCount != n-1 {
		return -1
	}

	return totalWeight
}

// find 查找元素所属的集合（路径压缩）
func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

// union 合并两个集合
func union(parent []int, x, y int) {
	rootX := find(parent, x)
	rootY := find(parent, y)
	if rootX != rootY {
		parent[rootY] = rootX
	}
}

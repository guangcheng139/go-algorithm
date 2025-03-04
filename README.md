# Go算法练习 - 面试高频题目集

这是一个专注于面试高频算法题目的Go语言实现集合，每个算法都配有详细注释和测试用例。

## 目录结构

```
go-algorithm/
├── basic/                       # 基础算法
│   ├── sort/                    # 排序算法
│   │   ├── bubble_sort.go       # 冒泡排序
│   │   ├── quick_sort.go        # 快速排序
│   │   ├── merge_sort.go        # 归并排序
│   │   └── sort_test.go         # 排序算法测试
│   └── search/                  # 查找算法
│       ├── binary_search.go     # 二分查找
│       └── search_test.go       # 查找算法测试
├── datastructure/               # 数据结构
│   ├── array/                   # 数组/切片
│   │   ├── array_problems.go    # 数组相关问题
│   │   └── array_test.go        # 数组问题测试
│   ├── string/                  # 字符串
│   │   ├── string_match.go      # 字符串匹配算法
│   │   └── string_test.go       # 字符串算法测试
│   ├── linkedlist/              # 链表
│   │   ├── linkedlist.go        # 链表实现
│   │   ├── problems.go          # 链表相关问题
│   │   └── linkedlist_test.go   # 链表测试
│   ├── stack/                   # 栈
│   │   ├── stack.go             # 栈实现
│   │   ├── problems.go          # 栈相关问题
│   │   └── stack_test.go        # 栈测试
│   ├── queue/                   # 队列
│   │   ├── queue.go             # 队列实现
│   │   ├── problems.go          # 队列相关问题
│   │   └── queue_test.go        # 队列测试
│   ├── tree/                    # 树
│   │   ├── binary_tree.go       # 二叉树
│   │   ├── bst.go               # 二叉搜索树
│   │   ├── problems.go          # 树相关问题
│   │   └── tree_test.go         # 树测试
│   └── graph/                   # 图
│       ├── graph.go             # 图的表示
│       ├── problems.go          # 图相关问题
│       └── graph_test.go        # 图测试
└── algorithm/                   # 算法题目（分类）
    ├── two_pointers/            # 双指针技巧
    │   ├── two_pointers.go      # 双指针相关题目
    │   └── two_pointers_test.go # 测试用例
    ├── sliding_window/          # 滑动窗口
    │   ├── sliding_window.go    # 滑动窗口相关题目
    │   └── sliding_window_test.go # 测试用例
    ├── binary_search/           # 二分搜索
    │   ├── binary_search.go     # 二分搜索相关题目
    │   └── binary_search_test.go # 测试用例
    ├── dfs_bfs/                 # 深度广度优先
    │   ├── dfs_bfs.go           # DFS/BFS相关题目
    │   └── dfs_bfs_test.go      # 测试用例
    ├── backtracking/            # 回溯算法
    │   ├── backtracking.go      # 回溯相关题目
    │   └── backtracking_test.go # 测试用例
    ├── dynamic_programming/     # 动态规划
    │   ├── dp.go                # DP相关题目
    │   └── dp_test.go           # 测试用例
    ├── greedy/                  # 贪心算法
    │   ├── greedy.go            # 贪心相关题目
    │   └── greedy_test.go       # 测试用例
    └── design/                  # 设计问题
        ├── lru_cache.go         # LRU缓存
        ├── lfu_cache.go         # LFU缓存
        └── design_test.go       # 测试用例
```

## 算法分类与题目

### 字符串算法

字符串匹配和处理是面试中常见的问题类型，掌握KMP、Rabin-Karp、Sunday等算法非常重要。

#### 字符串匹配示例

```go
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1。
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}

// 给定一个非空的字符串 s，检查它是否可以由它的一个子串重复多次构成。
func repeatedSubstringPattern(s string) bool {
	// 如果字符串长度小于2，则返回false
	if len(s) < 2 {
		return false
	}

	// 遍历字符串的一半
	for i := 0; i < len(s)/2; i++ {
		str := s[0 : i+1]
		if len(s)%len(str) == 0 {
			if s == strings.Repeat(str, len(s)/len(str)) {
				return true
			}
		}
	}
	return false
}
```

### 面试高频题目分类

下面列出各个分类中的高频面试题目：

#### 数组/切片

1. 两数之和
2. 三数之和
3. 合并两个有序数组
4. 寻找两个正序数组的中位数
5. 盛最多水的容器
6. 旋转数组
7. 移动零
8. 最大子数组和

#### 链表

1. 反转链表
2. 合并两个有序链表
3. 环形链表检测
4. 删除链表倒数第N个节点
5. 链表中间节点
6. 相交链表
7. 回文链表判断

#### 字符串

1. 字符串匹配（KMP、Sunday等）
2. 最长回文子串
3. 有效的括号
4. 最长公共前缀
5. 字符串转整数(atoi)
6. 验证回文串

#### 栈与队列

1. 有效的括号
2. 最小栈
3. 逆波兰表达式求值
4. 队列实现栈/栈实现队列
5. 滑动窗口最大值

#### 树

1. 二叉树的前/中/后序遍历
2. 二叉树的层序遍历
3. 验证二叉搜索树
4. 对称二叉树
5. 二叉树的最大深度
6. 二叉树的最近公共祖先
7. 二叉树的序列化与反序列化

#### 图

1. 岛屿数量
2. 课程表（拓扑排序）
3. 单词接龙
4. 克隆图
5. 最短路径算法

#### 排序与搜索

1. 各种排序算法实现与分析
2. 第k个最大元素
3. 搜索旋转排序数组
4. 寻找峰值
5. 二分查找变种

#### 动态规划

1. 爬楼梯
2. 打家劫舍
3. 买卖股票的最佳时机
4. 最长递增子序列
5. 最长公共子序列
6. 0-1背包问题
7. 编辑距离
8. 正则表达式匹配

#### 贪心算法

1. 跳跃游戏
2. 加油站
3. 分发糖果
4. 任务调度器

#### 设计问题

1. LRU缓存
2. LFU缓存
3. 数据流的中位数
4. 设计推特

## 测试方法

每个算法都配有相应的测试用例，可以通过以下方式运行测试：

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./datastructure/string/

# 运行特定测试用例
go test -run TestStrStr ./datastructure/string/
```

## 学习建议

1. 按照分类系统学习，先掌握基础数据结构
2. 针对每种算法类型，先理解基本思想再解题
3. 多练习变种题目，培养举一反三的能力
4. 注重时间和空间复杂度分析
5. 定期复习，防止遗忘

## 参考资源

- [algorithm-pattern](https://github.com/greyireland/algorithm-pattern)
- LeetCode题库
- 《算法》(第4版)
- 《剑指Offer》

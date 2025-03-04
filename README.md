# Golang算法练习计划目录结构

这是一个系统化的Golang算法练习计划，采用模块化设计，帮助你从基础到高级逐步掌握算法与数据结构。

## 第一阶段：Golang基础与简单数据结构

### 1.1 Golang基础复习
- 变量、类型和控制结构
- 函数、指针和结构体
- 切片、映射和数组的基本操作

**练习题：**
1. 实现一个函数，将整数切片中的所有元素翻倍
2. 创建一个学生成绩管理系统，使用结构体和映射
3. 实现一个简单的命令行计算器
4. 编写一个函数，合并两个有序数组

### 1.2 简单数据结构实现
- 数组与切片操作
- 链表（单链表、双链表、循环链表）
- 栈与队列
- 哈希表

**练习题：**
1. 实现一个单链表，包含插入、删除和查找操作
2. 使用切片实现一个栈，包含push、pop和peek操作
3. 使用链表实现一个队列，包含enqueue和dequeue操作
4. 实现一个简单的哈希表，处理冲突
5. 设计一个循环队列

## 第二阶段：基本算法思想

### 2.1 时间复杂度与空间复杂度分析
- Big O表示法
- 常见算法复杂度分析

**练习题：**
1. 分析冒泡排序的时间复杂度
2. 分析二分查找的时间复杂度
3. 比较不同排序算法的性能
4. 优化一个O(n²)的算法到O(n log n)

### 2.2 排序算法
- 冒泡排序、选择排序、插入排序
- 快速排序、归并排序、堆排序
- 计数排序、桶排序、基数排序

**练习题：**
1. 实现冒泡排序和插入排序，比较性能
2. 实现快速排序，处理不同的pivot选择策略
3. 实现归并排序，并应用到链表排序
4. 对包含大量重复元素的数组进行排序，比较不同算法性能
5. 实现基数排序处理非负整数数组

### 2.3 查找算法
- 顺序查找
- 二分查找
- 哈希查找

**练习题：**
1. 实现二分查找（递归和迭代两种方式）
2. 在旋转排序数组中查找目标值
3. 实现查找第一个和最后一个出现的位置
4. 设计一个数据结构，支持O(1)时间复杂度的查找最小值操作

## 第三阶段：常见算法问题

### 3.1 递归与分治
- 递归基础问题（阶乘、斐波那契）
- 分治算法（归并排序、快速排序）
- 回溯算法（N皇后、数独求解）

**练习题：**
1. 使用递归和动态规划两种方式计算斐波那契数列
2. 实现汉诺塔问题的求解
3. 使用回溯算法解决八皇后问题
4. 实现数独求解器
5. 生成所有可能的括号组合

### 3.2 贪心算法
- 活动选择问题
- 哈夫曼编码
- 最小生成树（Prim、Kruskal）

**练习题：**
1. 解决会议室安排问题
2. 实现哈夫曼编码和解码
3. 使用贪心算法解决背包问题的近似解
4. 实现Prim算法求最小生成树
5. 解决零钱兑换问题

### 3.3 动态规划
- 基本概念（最优子结构、重叠子问题）
- 经典问题（背包问题、最长公共子序列）
- 路径问题（最短路径、矩阵路径）

**练习题：**
1. 实现0-1背包问题
2. 计算最长公共子序列
3. 计算最长递增子序列
4. 矩阵链乘法最优解
5. 编辑距离问题

## 第四阶段：高级数据结构

### 4.1 树结构
- 二叉树、二叉搜索树
- 平衡树（AVL树、红黑树）
- 堆（最大堆、最小堆）
- 并查集

**练习题：**
1. 实现二叉搜索树的插入、删除和查找
2. 实现二叉树的前序、中序和后序遍历
3. 实现AVL树的自平衡操作
4. 创建最小堆并实现堆排序
5. 使用并查集解决朋友圈问题

### 4.2 图算法
- 图的表示（邻接矩阵、邻接表）
- 图的遍历（BFS、DFS）
- 最短路径（Dijkstra、Bellman-Ford、Floyd）
- 拓扑排序

**练习题：**
1. 实现图的BFS和DFS遍历
2. 判断图是否为二分图
3. 实现Dijkstra算法求单源最短路径
4. 实现拓扑排序解决课程安排问题
5. 检测图中是否存在环

### 4.3 高级数据结构
- Trie树（前缀树）
- 线段树
- 树状数组
- 跳表

**练习题：**
1. 实现Trie树并解决单词搜索问题
2. 使用线段树解决区间查询问题
3. 实现树状数组求前缀和
4. 实现跳表的基本操作
5. 使用Trie树实现自动补全功能

## 第五阶段：实战与优化

### 5.1 算法设计技巧
- 双指针技巧
- 滑动窗口
- 位运算优化
- 前缀和与差分

**练习题：**
1. 使用双指针解决三数之和问题
2. 实现滑动窗口找出最长无重复字符子串
3. 使用位运算实现不用加减乘除做加法
4. 使用前缀和解决子数组和问题
5. 实现差分数组处理区间更新

### 5.2 经典算法问题集
- LeetCode热门150题
- 面试常见算法题

**推荐题目：**
1. 反转链表
2. 合并两个有序链表
3. LRU缓存设计
4. 有效的括号
5. 岛屿数量

### 5.3 实战项目
- 简单搜索引擎
- 缓存系统设计（LRU/LFU）
- 简易数据库索引实现

**项目练习：**
1. 实现一个简单的基于Trie的搜索引擎
2. 设计LRU和LFU缓存，对比性能差异
3. 实现一个简单的B树索引
4. 开发一个简易的HTTP路由器
5. 实现一个基于令牌桶的限流器

## 学习资源

### 推荐书籍
- 《数据结构与算法分析：Go语言描述》
- 《算法（第4版）》
- 《编程珠玑》

### 在线资源
- LeetCode、力扣中文网
- GitHub上的Go算法仓库
- Golang官方文档

## 学习建议
1. 每天保持持续学习，至少1-2小时
2. 先理解算法思想，再尝试独立实现
3. 多做练习题，巩固理解
4. 复习已学内容，防止遗忘
5. 尝试优化你的解决方案

## 目录结构建议

golang-algorithm-practice/
├── 01-basics/
│   ├── 01-go-basics/
│   │   ├── 01-double-elements.go         # 切片元素翻倍函数
│   │   ├── 02-student-management.go      # 学生成绩管理系统
│   │   ├── 03-simple-calculator.go       # 简单命令行计算器
│   │   └── 04-merge-sorted-arrays.go     # 合并有序数组
│   └── 02-simple-data-structures/
│       ├── 01-linked-list/
│       │   ├── singly_linked_list.go     # 单链表实现
│       │   ├── doubly_linked_list.go     # 双链表实现
│       │   └── circular_linked_list.go   # 循环链表实现
│       ├── 02-stack/
│       │   ├── slice_stack.go            # 基于切片的栈实现
│       │   └── linked_stack.go           # 基于链表的栈实现
│       ├── 03-queue/
│       │   ├── linked_queue.go           # 基于链表的队列
│       │   └── circular_queue.go         # 循环队列
│       └── 04-hash-table/
│           └── hash_table.go             # 哈希表实现
├── 02-basic-algorithms/
│   ├── 01-complexity/
│   │   ├── bubble_sort_analysis.go       # 冒泡排序复杂度分析
│   │   ├── binary_search_analysis.go     # 二分查找复杂度分析
│   │   └── algorithm_comparison.go       # 不同算法性能比较
│   ├── 02-sorting/
│   │   ├── bubble_sort.go                # 冒泡排序
│   │   ├── selection_sort.go             # 选择排序
│   │   ├── insertion_sort.go             # 插入排序
│   │   ├── quick_sort.go                 # 快速排序
│   │   ├── merge_sort.go                 # 归并排序
│   │   ├── heap_sort.go                  # 堆排序
│   │   ├── counting_sort.go              # 计数排序
│   │   ├── bucket_sort.go                # 桶排序
│   │   └── radix_sort.go                 # 基数排序
│   └── 03-searching/
│       ├── sequential_search.go          # 顺序查找
│       ├── binary_search.go              # 二分查找
│       ├── rotated_array_search.go       # 旋转数组中查找
│       └── hash_search.go                # 哈希查找
├── 03-common-problems/
│   ├── 01-recursion-divide-conquer/
│   │   ├── fibonacci.go                  # 斐波那契数列
│   │   ├── hanoi_tower.go                # 汉诺塔问题
│   │   ├── n_queens.go                   # 八皇后问题
│   │   ├── sudoku_solver.go              # 数独求解器
│   │   └── generate_parentheses.go       # 生成括号组合
│   ├── 02-greedy/
│   │   ├── meeting_rooms.go              # 会议室安排问题
│   │   ├── huffman_coding.go             # 哈夫曼编码
│   │   ├── knapsack_greedy.go            # 贪心背包问题
│   │   ├── prim_algorithm.go             # Prim最小生成树
│   │   └── coin_change_greedy.go         # 零钱兑换问题
│   └── 03-dynamic-programming/
│       ├── knapsack_01.go                # 0-1背包问题
│       ├── longest_common_subsequence.go # 最长公共子序列
│       ├── longest_increasing_subsequence.go # 最长递增子序列
│       ├── matrix_chain_multiplication.go # 矩阵链乘法
│       └── edit_distance.go              # 编辑距离问题
├── 04-advanced-data-structures/
│   ├── 01-trees/
│   │   ├── binary_search_tree.go         # 二叉搜索树
│   │   ├── tree_traversal.go             # 树的遍历
│   │   ├── avl_tree.go                   # AVL树
│   │   ├── min_heap.go                   # 最小堆
│   │   └── union_find.go                 # 并查集
│   ├── 02-graphs/
│   │   ├── graph_representation/
│   │   │   ├── adjacency_matrix.go       # 邻接矩阵表示
│   │   │   └── adjacency_list.go         # 邻接表表示
│   │   ├── graph_traversal.go            # 图的BFS和DFS遍历
│   │   ├── bipartite_graph.go            # 二分图判断
│   │   ├── dijkstra.go                   # Dijkstra最短路径
│   │   ├── bellman_ford.go               # Bellman-Ford算法
│   │   ├── floyd_warshall.go             # Floyd-Warshall算法
│   │   ├── topological_sort.go           # 拓扑排序
│   │   └── cycle_detection.go            # 环检测
│   └── 03-advanced-structures/
│       ├── trie.go                       # Trie树(前缀树)
│       ├── segment_tree.go               # 线段树
│       ├── binary_indexed_tree.go        # 树状数组
│       └── skip_list.go                  # 跳表
├── 05-practice/
│   ├── 01-algorithm-techniques/
│   │   ├── two_pointers.go               # 双指针技巧
│   │   ├── sliding_window.go             # 滑动窗口
│   │   ├── bit_manipulation.go           # 位运算优化
│   │   ├── prefix_sum.go                 # 前缀和
│   │   └── difference_array.go           # 差分数组
│   ├── 02-classic-problems/
│   │   ├── reverse_linked_list.go        # 反转链表
│   │   ├── merge_sorted_lists.go         # 合并有序链表
│   │   ├── lru_cache.go                  # LRU缓存
│   │   ├── valid_parentheses.go          # 有效的括号
│   │   └── number_of_islands.go          # 岛屿数量
│   └── 03-projects/
│       ├── trie_search_engine/           # 基于Trie的搜索引擎
│       ├── cache_system/                 # LRU和LFU缓存系统
│       │   ├── lru_cache.go
│       │   └── lfu_cache.go
│       ├── b_tree_index/                 # B树索引实现
│       ├── http_router/                  # HTTP路由器
│       └── rate_limiter/                 # 令牌桶限流器
└── README.md

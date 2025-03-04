package design

// LRUCache 是一个最近最少使用（Least Recently Used）缓存结构
// 当缓存达到容量上限时，它会删除最久未使用的项目
type LRUCache struct {
	capacity int              // 缓存容量
	cache    map[int]*DLLNode // 存储键到节点的映射
	head     *DLLNode         // 双向链表头节点（最近使用的）
	tail     *DLLNode         // 双向链表尾节点（最久未使用的）
}

// DLLNode 双向链表节点
type DLLNode struct {
	key   int
	value int
	prev  *DLLNode
	next  *DLLNode
}

// Constructor 初始化LRU缓存
// 时间复杂度: O(1)
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLLNode, capacity),
	}

	// 创建哑头和哑尾，简化边界情况处理
	lru.head = &DLLNode{}
	lru.tail = &DLLNode{}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

// Get 获取键对应的值，如果不存在返回-1
// 时间复杂度: O(1)
func (lru *LRUCache) Get(key int) int {
	node, exists := lru.cache[key]
	if !exists {
		return -1
	}

	// 将节点移至链表头部（表示最近使用）
	lru.moveToHead(node)
	return node.value
}

// Put 插入键值对，如果达到容量，删除最久未使用的项
// 时间复杂度: O(1)
func (lru *LRUCache) Put(key int, value int) {
	// 如果容量为0，直接返回
	if lru.capacity <= 0 {
		return
	}

	// 检查键是否已存在
	if node, exists := lru.cache[key]; exists {
		// 更新值并移动到链表头部
		node.value = value
		lru.moveToHead(node)
		return
	}

	// 如果达到容量，删除最久未使用的节点（链表尾部）
	if len(lru.cache) >= lru.capacity {
		// 获取和删除尾部节点
		tail := lru.tail.prev
		lru.removeNode(tail)
		delete(lru.cache, tail.key)
	}

	// 创建新节点并添加到链表头部
	newNode := &DLLNode{
		key:   key,
		value: value,
	}
	lru.addToHead(newNode)
	lru.cache[key] = newNode
}

// 将节点移至链表头部
func (lru *LRUCache) moveToHead(node *DLLNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// 从链表中删除节点
func (lru *LRUCache) removeNode(node *DLLNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 添加节点到链表头部
func (lru *LRUCache) addToHead(node *DLLNode) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

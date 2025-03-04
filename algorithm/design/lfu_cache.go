package design

// LFUCache 是一个最不经常使用（Least Frequently Used）缓存结构
// 当缓存达到容量上限时，它会删除使用频率最低的项目
// 如果有多个使用频率相同的项目，则删除最久未使用的项目
type LFUCache struct {
	capacity int                 // 缓存容量
	minFreq  int                 // 当前最小频率
	cache    map[int]*LFUNode    // 键到节点的映射
	freqMap  map[int]*DoublyList // 频率到双向链表的映射
}

// LFUNode LFU节点，存储键值对和使用频率
type LFUNode struct {
	key        int
	value      int
	freq       int
	prev, next *LFUNode
}

// DoublyList 双向链表，用于存储相同频率的节点
type DoublyList struct {
	head, tail *LFUNode
	size       int
}

// NewDoublyList 创建新的双向链表
func NewDoublyList() *DoublyList {
	head := &LFUNode{}
	tail := &LFUNode{}
	head.next = tail
	tail.prev = head

	return &DoublyList{
		head: head,
		tail: tail,
		size: 0,
	}
}

// AddNode 添加节点到链表头部
func (dl *DoublyList) AddNode(node *LFUNode) {
	node.next = dl.head.next
	node.prev = dl.head
	dl.head.next.prev = node
	dl.head.next = node
	dl.size++
}

// RemoveNode 从链表中删除节点
func (dl *DoublyList) RemoveNode(node *LFUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	dl.size--
}

// RemoveTail 移除并返回链表尾部节点
func (dl *DoublyList) RemoveTail() *LFUNode {
	if dl.size == 0 {
		return nil
	}

	node := dl.tail.prev
	dl.RemoveNode(node)
	return node
}

// NewLFUCache 创建新的LFU缓存
// 时间复杂度: O(1)
func NewLFUCache(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*LFUNode),
		freqMap:  make(map[int]*DoublyList),
	}
}

// Get 获取键对应的值，如果不存在返回-1
// 时间复杂度: O(1)
func (lfu *LFUCache) Get(key int) int {
	if lfu.capacity == 0 {
		return -1
	}

	node, exists := lfu.cache[key]
	if !exists {
		return -1
	}

	// 增加节点频率并移动到相应频率的链表
	lfu.updateFreq(node)
	return node.value
}

// Put 插入键值对，如果达到容量，删除使用频率最低的项
// 时间复杂度: O(1)
func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}

	// 如果键已存在，更新值和频率
	if node, exists := lfu.cache[key]; exists {
		node.value = value
		lfu.updateFreq(node)
		return
	}

	// 如果达到容量上限，删除使用频率最低的项
	if len(lfu.cache) >= lfu.capacity {
		// 获取最小频率链表
		list := lfu.freqMap[lfu.minFreq]
		// 删除最小频率链表中的最后一个节点（最久未使用）
		toDelete := list.RemoveTail()
		// 从缓存中删除
		delete(lfu.cache, toDelete.key)

		// 如果链表为空，删除频率映射
		if list.size == 0 {
			delete(lfu.freqMap, lfu.minFreq)
		}
	}

	// 创建新节点，频率为1
	newNode := &LFUNode{
		key:   key,
		value: value,
		freq:  1,
	}

	// 添加到频率为1的链表
	if _, exists := lfu.freqMap[1]; !exists {
		lfu.freqMap[1] = NewDoublyList()
	}

	lfu.freqMap[1].AddNode(newNode)
	lfu.cache[key] = newNode
	lfu.minFreq = 1 // 新节点的频率为1，设为最小频率
}

// updateFreq 增加节点的使用频率并移动到相应频率的链表
func (lfu *LFUCache) updateFreq(node *LFUNode) {
	// 从当前频率链表中删除节点
	oldFreq := node.freq
	list := lfu.freqMap[oldFreq]
	list.RemoveNode(node)

	// 如果链表为空且是最小频率，更新最小频率
	if list.size == 0 && oldFreq == lfu.minFreq {
		delete(lfu.freqMap, oldFreq)
		lfu.minFreq++
	} else if list.size == 0 {
		delete(lfu.freqMap, oldFreq)
	}

	// 增加节点频率
	node.freq++
	newFreq := node.freq

	// 添加到新频率的链表
	if _, exists := lfu.freqMap[newFreq]; !exists {
		lfu.freqMap[newFreq] = NewDoublyList()
	}

	lfu.freqMap[newFreq].AddNode(node)
}

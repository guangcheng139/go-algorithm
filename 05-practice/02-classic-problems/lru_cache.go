package classic

// LRUNode 表示LRU缓存中的一个节点
type LRUNode struct {
	Key   int
	Value int
	Prev  *LRUNode
	Next  *LRUNode
}

// LRUCache 表示LRU（最近最少使用）缓存
type LRUCache struct {
	Capacity int              // 缓存容量
	Size     int              // 当前大小
	Cache    map[int]*LRUNode // 哈希表，用于O(1)时间查找
	Head     *LRUNode         // 双向链表头（最近使用）
	Tail     *LRUNode         // 双向链表尾（最久未使用）
}

// NewLRUCache 创建一个新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	// 创建哑头和哑尾节点，简化边界情况处理
	head := &LRUNode{0, 0, nil, nil}
	tail := &LRUNode{0, 0, nil, nil}
	head.Next = tail
	tail.Prev = head

	return &LRUCache{
		Capacity: capacity,
		Size:     0,
		Cache:    make(map[int]*LRUNode),
		Head:     head,
		Tail:     tail,
	}
}

// Get 获取缓存中的值
// 如果键不存在，返回-1
func (c *LRUCache) Get(key int) int {
	node, exists := c.Cache[key]
	if !exists {
		return -1
	}

	// 将节点移到链表头部（表示最近使用）
	c.moveToHead(node)
	return node.Value
}

// Put 将键值对放入缓存
// 如果键已存在，更新值
// 如果缓存已满，删除最久未使用的项
func (c *LRUCache) Put(key int, value int) {
	// 检查键是否已存在
	node, exists := c.Cache[key]
	if exists {
		// 更新值并移到链表头部
		node.Value = value
		c.moveToHead(node)
		return
	}

	// 创建新节点
	newNode := &LRUNode{
		Key:   key,
		Value: value,
	}

	// 添加到哈希表和链表头部
	c.Cache[key] = newNode
	c.addToHead(newNode)
	c.Size++

	// 如果超出容量，删除最久未使用的节点（链表尾部）
	if c.Size > c.Capacity {
		removed := c.removeTail()
		delete(c.Cache, removed.Key)
		c.Size--
	}
}

// addToHead 将节点添加到链表头部
func (c *LRUCache) addToHead(node *LRUNode) {
	node.Prev = c.Head
	node.Next = c.Head.Next
	c.Head.Next.Prev = node
	c.Head.Next = node
}

// removeNode 从链表中删除节点
func (c *LRUCache) removeNode(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// moveToHead 将节点移动到链表头部
func (c *LRUCache) moveToHead(node *LRUNode) {
	c.removeNode(node)
	c.addToHead(node)
}

// removeTail 删除链表尾部节点并返回
func (c *LRUCache) removeTail() *LRUNode {
	node := c.Tail.Prev
	c.removeNode(node)
	return node
}

// Size 返回当前缓存中的项数
func (c *LRUCache) GetSize() int {
	return c.Size
}

// Clear 清空缓存
func (c *LRUCache) Clear() {
	c.Head.Next = c.Tail
	c.Tail.Prev = c.Head
	c.Cache = make(map[int]*LRUNode)
	c.Size = 0
}

// GetKeys 获取缓存中的所有键（按最近使用顺序）
func (c *LRUCache) GetKeys() []int {
	keys := make([]int, 0, c.Size)
	current := c.Head.Next
	for current != c.Tail {
		keys = append(keys, current.Key)
		current = current.Next
	}
	return keys
}

# 设计问题模块

本模块包含了常见的算法设计问题的实现，重点关注数据结构和算法的设计与优化。

## 实现的缓存结构

### LRU缓存 (Least Recently Used)

LRU缓存保留最近使用的数据项，当缓存超过容量限制时，会删除最久未使用的项目。

实现原理：
- 使用哈希表 + 双向链表的组合数据结构
- 哈希表提供O(1)的查找时间
- 双向链表用于维护使用顺序，便于O(1)时间删除和添加节点
- 最近使用的节点放在链表头部，最久未使用的节点位于链表尾部

应用场景：
- 页面置换算法
- 数据库缓存
- 网页缓存

### LFU缓存 (Least Frequently Used)

LFU缓存保留最常使用的数据项，当缓存超过容量限制时，会删除使用频率最低的项目。如果多个项目具有相同的使用频率，则删除最久未使用的项目。

实现原理：
- 使用两个哈希表和多个双向链表
- 第一个哈希表用于O(1)时间根据键查找节点
- 第二个哈希表将频率映射到相应的双向链表，每个链表包含具有相同频率的节点
- 访问节点时增加其频率并移动到相应频率的链表
- 保持最小频率的记录，便于快速找到要删除的节点

应用场景：
- 内容分发网络(CDN)
- 数据库查询缓存
- 操作系统页面置换

## 性能对比

| 缓存策略 | 优点 | 缺点 | 适用场景 |
|---------|------|------|----------|
| LRU | 实现简单，适合时间局部性强的应用 | 对突发性访问模式不友好 | 大多数常规缓存场景 |
| LFU | 可以更好地反映长期访问模式 | 实现复杂，对冷启动不友好 | 长时间运行且访问模式相对稳定的系统 |

## 使用示例

```go
// LRU缓存示例
lru := Constructor(2)  // 创建容量为2的LRU缓存
lru.Put(1, 1)  // 添加键值对 (1,1)
lru.Put(2, 2)  // 添加键值对 (2,2)
lru.Get(1)     // 返回 1
lru.Put(3, 3)  // 移除键 2，添加键 3
lru.Get(2)     // 返回 -1 (未找到)
lru.Put(4, 4)  // 移除键 1，添加键 4
lru.Get(1)     // 返回 -1 (未找到)
lru.Get(3)     // 返回 3
lru.Get(4)     // 返回 4

// LFU缓存示例
lfu := NewLFUCache(2)  // 创建容量为2的LFU缓存
lfu.Put(1, 1)  // 添加键值对 (1,1)
lfu.Put(2, 2)  // 添加键值对 (2,2)
lfu.Get(1)     // 返回 1，键1的频率增加到2
lfu.Put(3, 3)  // 移除键 2 (频率为1)，添加键 3
lfu.Get(2)     // 返回 -1 (未找到)
lfu.Get(3)     // 返回 3，键3的频率增加到2
lfu.Put(4, 4)  // 移除键 1 (最久未使用的频率为2的键)，添加键 4
lfu.Get(1)     // 返回 -1 (未找到)
``` 
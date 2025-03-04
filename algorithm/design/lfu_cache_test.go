package design

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		lfu := NewLFUCache(2)

		// 添加两个元素
		lfu.Put(1, 1)
		lfu.Put(2, 2)

		// 验证可以获取元素
		if val := lfu.Get(1); val != 1 {
			t.Errorf("期望 Get(1) = 1, 实际得到 %d", val)
		}

		// 添加第三个元素，应该移除使用频率最低的元素 (键2)
		lfu.Put(3, 3)

		// 验证键2已被移除
		if val := lfu.Get(2); val != -1 {
			t.Errorf("期望 Get(2) = -1, 实际得到 %d", val)
		}

		// 验证键1和键3仍存在
		if val := lfu.Get(1); val != 1 {
			t.Errorf("期望 Get(1) = 1, 实际得到 %d", val)
		}
		if val := lfu.Get(3); val != 3 {
			t.Errorf("期望 Get(3) = 3, 实际得到 %d", val)
		}
	})

	t.Run("频率相同时移除最久未使用的元素", func(t *testing.T) {
		lfu := NewLFUCache(2)

		// 添加两个元素
		lfu.Put(1, 1)
		lfu.Put(2, 2)

		// 访问两个元素，使其频率都为2
		lfu.Get(1)
		lfu.Get(2)

		// 再次访问键1，使其频率为3
		lfu.Get(1)

		// 添加第三个元素，应该移除使用频率较低的元素 (键2)
		lfu.Put(3, 3)

		// 验证键2已被移除，键1和键3仍存在
		if val := lfu.Get(2); val != -1 {
			t.Errorf("期望 Get(2) = -1, 实际得到 %d", val)
		}
		if val := lfu.Get(1); val != 1 {
			t.Errorf("期望 Get(1) = 1, 实际得到 %d", val)
		}
		if val := lfu.Get(3); val != 3 {
			t.Errorf("期望 Get(3) = 3, 实际得到 %d", val)
		}
	})

	t.Run("更新已存在的元素", func(t *testing.T) {
		lfu := NewLFUCache(2)

		// 添加两个元素
		lfu.Put(1, 1)
		lfu.Put(2, 2)

		// 更新键1的值
		lfu.Put(1, 10)

		// 验证键1的值已更新，且频率增加
		if val := lfu.Get(1); val != 10 {
			t.Errorf("期望 Get(1) = 10, 实际得到 %d", val)
		}

		// 添加第三个元素，应该移除使用频率最低的元素 (键2)
		lfu.Put(3, 3)

		// 验证键2已被移除，键1和键3仍存在
		if val := lfu.Get(2); val != -1 {
			t.Errorf("期望 Get(2) = -1, 实际得到 %d", val)
		}
		if val := lfu.Get(1); val != 10 {
			t.Errorf("期望 Get(1) = 10, 实际得到 %d", val)
		}
		if val := lfu.Get(3); val != 3 {
			t.Errorf("期望 Get(3) = 3, 实际得到 %d", val)
		}
	})

	t.Run("获取非存在的元素", func(t *testing.T) {
		lfu := NewLFUCache(2)

		// 获取不存在的元素应返回-1
		if val := lfu.Get(1); val != -1 {
			t.Errorf("期望 Get(1) = -1, 实际得到 %d", val)
		}
	})

	t.Run("零容量缓存", func(t *testing.T) {
		lfu := NewLFUCache(0)

		// 添加元素到零容量缓存
		lfu.Put(1, 1)

		// 验证元素未被添加
		if val := lfu.Get(1); val != -1 {
			t.Errorf("期望 Get(1) = -1, 实际得到 %d", val)
		}
	})

	t.Run("复杂操作序列", func(t *testing.T) {
		lfu := NewLFUCache(3)

		// 添加三个元素
		lfu.Put(1, 1) // freq: 1
		lfu.Put(2, 2) // freq: 1
		lfu.Put(3, 3) // freq: 1

		// 访问键1和键2，使其频率为2
		lfu.Get(1) // freq: 2
		lfu.Get(2) // freq: 2

		// 添加第四个元素，应该移除使用频率最低的元素 (键3)
		lfu.Put(4, 4) // freq: 1

		// 验证键3已被移除
		if val := lfu.Get(3); val != -1 {
			t.Errorf("期望 Get(3) = -1, 实际得到 %d", val)
		}

		// 访问键1和键4，使键1频率为3，键4频率为2
		lfu.Get(1) // freq: 3
		lfu.Get(4) // freq: 2

		// 添加第五个元素，应该移除使用频率最低的元素 (键2)
		lfu.Put(5, 5) // freq: 1

		// 验证键2已被移除
		if val := lfu.Get(2); val != -1 {
			t.Errorf("期望 Get(2) = -1, 实际得到 %d", val)
		}

		// 验证其他元素仍存在
		if val := lfu.Get(1); val != 1 {
			t.Errorf("期望 Get(1) = 1, 实际得到 %d", val)
		}
		if val := lfu.Get(4); val != 4 {
			t.Errorf("期望 Get(4) = 4, 实际得到 %d", val)
		}
		if val := lfu.Get(5); val != 5 {
			t.Errorf("期望 Get(5) = 5, 实际得到 %d", val)
		}
	})
}

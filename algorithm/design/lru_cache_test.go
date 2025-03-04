package design

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		lru := Constructor(2)

		// 添加两个元素
		lru.Put(1, 1)
		lru.Put(2, 2)

		// 验证可以获取元素
		if val := lru.Get(1); val != 1 {
			t.Errorf("期望 Get(1) = 1, 实际得到 %d", val)
		}

		// 添加第三个元素，应该移除最少使用的元素 (键2)
		lru.Put(3, 3)

		// 验证键2已被移除
		if val := lru.Get(2); val != -1 {
			t.Errorf("期望 Get(2) = -1, 实际得到 %d", val)
		}

		// 再添加一个元素，应该移除最少使用的元素 (键1)
		lru.Put(4, 4)

		// 验证键1已被移除，键3和键4仍存在
		if val := lru.Get(1); val != -1 {
			t.Errorf("期望 Get(1) = -1, 实际得到 %d", val)
		}
		if val := lru.Get(3); val != 3 {
			t.Errorf("期望 Get(3) = 3, 实际得到 %d", val)
		}
		if val := lru.Get(4); val != 4 {
			t.Errorf("期望 Get(4) = 4, 实际得到 %d", val)
		}
	})

	t.Run("更新已存在的元素", func(t *testing.T) {
		lru := Constructor(2)

		// 添加两个元素
		lru.Put(1, 1)
		lru.Put(2, 2)

		// 更新键1的值
		lru.Put(1, 10)

		// 验证键1的值已更新
		if val := lru.Get(1); val != 10 {
			t.Errorf("期望 Get(1) = 10, 实际得到 %d", val)
		}

		// 添加新元素，应该移除最少使用的元素 (键2)
		lru.Put(3, 3)

		// 验证键2已被移除，键1和键3仍存在
		if val := lru.Get(2); val != -1 {
			t.Errorf("期望 Get(2) = -1, 实际得到 %d", val)
		}
		if val := lru.Get(1); val != 10 {
			t.Errorf("期望 Get(1) = 10, 实际得到 %d", val)
		}
		if val := lru.Get(3); val != 3 {
			t.Errorf("期望 Get(3) = 3, 实际得到 %d", val)
		}
	})

	t.Run("获取非存在的元素", func(t *testing.T) {
		lru := Constructor(2)

		// 获取不存在的元素应返回-1
		if val := lru.Get(1); val != -1 {
			t.Errorf("期望 Get(1) = -1, 实际得到 %d", val)
		}
	})

	t.Run("零容量缓存", func(t *testing.T) {
		lru := Constructor(0)

		// 添加元素到零容量缓存
		lru.Put(1, 1)

		// 验证元素未被添加
		if val := lru.Get(1); val != -1 {
			t.Errorf("期望 Get(1) = -1, 实际得到 %d", val)
		}
	})

	t.Run("顺序验证", func(t *testing.T) {
		lru := Constructor(3)

		// 添加三个元素
		lru.Put(1, 1)
		lru.Put(2, 2)
		lru.Put(3, 3)

		// 访问键1，使其成为最近使用
		if val := lru.Get(1); val != 1 {
			t.Errorf("期望 Get(1) = 1, 实际得到 %d", val)
		}

		// 添加第四个元素，应该移除最少使用的元素 (键2)
		lru.Put(4, 4)

		// 验证键2已被移除，键1、键3和键4仍存在
		if val := lru.Get(2); val != -1 {
			t.Errorf("期望 Get(2) = -1, 实际得到 %d", val)
		}
		if val := lru.Get(1); val != 1 {
			t.Errorf("期望 Get(1) = 1, 实际得到 %d", val)
		}
		if val := lru.Get(3); val != 3 {
			t.Errorf("期望 Get(3) = 3, 实际得到 %d", val)
		}
		if val := lru.Get(4); val != 4 {
			t.Errorf("期望 Get(4) = 4, 实际得到 %d", val)
		}
	})
}

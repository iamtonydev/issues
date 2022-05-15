package lru_cache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       *sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mu:       new(sync.Mutex),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheObj := cacheItem{
		key:   key,
		value: value,
	}
	listItem, find := c.items[key]

	if find {
		listItem.Value = cacheObj
		c.queue.MoveToFront(listItem)
	} else {
		if c.queue.Len() == c.capacity {
			lastCacheItem := c.queue.Back()
			c.queue.Remove(lastCacheItem)
			delete(c.items, lastCacheItem.Value.(cacheItem).key)
		}
		listItem = c.queue.PushFront(cacheObj)
		c.items[key] = listItem
	}
	return find
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	listItem, find := c.items[key]
	if find {
		c.queue.MoveToFront(listItem)
		return listItem.Value.(cacheItem).value, find
	}
	return nil, find
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

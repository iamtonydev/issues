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
	mu       sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    newList(),
		items:    make(map[Key]*ListItem, capacity),
		mu:       sync.Mutex{},
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := cacheItem{
		key:   key,
		value: value,
	}

	if listItem, found := c.items[key]; found {
		listItem.Value = item
		c.queue.MoveToFront(listItem)
		c.items[key] = listItem
		return found
	}

	if c.queue.Len() == c.capacity {
		lastCacheItem := c.queue.Back()
		c.queue.Remove(lastCacheItem)
		delete(c.items, lastCacheItem.Value.(cacheItem).key)
	}
	listItem := c.queue.PushFront(item)
	c.items[key] = listItem

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	listItem, found := c.items[key]
	if found {
		c.queue.MoveToFront(listItem)
		return listItem.Value.(cacheItem).value, found
	}
	return nil, found
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = newList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

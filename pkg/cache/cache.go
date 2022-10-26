package cache

import "container/list"

type Cache struct {
	cap int
	l   *list.List
	m   map[string]*list.Element
}

type CacheData struct {
	key   string
	value interface{}
}

type CacheResp struct {
	Found bool
	Data  interface{}
}

func NewCache(capacity int) *Cache {
	return &Cache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[string]*list.Element, capacity),
	}
}

func (c *Cache) Get(key string) CacheResp {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		val := node.Value.(*list.Element).Value.(CacheData).value
		// move node to front
		c.l.MoveToFront(node)
		return CacheResp{
			Found: true,
			Data:  val,
		}
	}
	return CacheResp{
		Found: false,
	}
}

func (c *Cache) Put(key string, value interface{}) {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		// move the node to front
		c.l.MoveToFront(node)
		// update the value of a list node
		node.Value.(*list.Element).Value = CacheData{key: key, value: value}
	} else {
		// delete the last list node if the list is full
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			idx := c.l.Back().Value.(*list.Element).Value.(CacheData).key
			// delete the node pointer in the hash map by key
			delete(c.m, idx)
			// remove the last list node
			c.l.Remove(c.l.Back())
		}
		// initialize a list node
		node := &list.Element{
			Value: CacheData{
				key:   key,
				value: value,
			},
		}
		// push the new list node into the list
		ptr := c.l.PushFront(node)
		// save the node pointer in the hash map
		c.m[key] = ptr
	}
}

func (c *Cache) Remove(key string) {
	if node, ok := c.m[key]; ok {
		// get the key that we want to delete
		idx := node.Value.(*list.Element).Value.(CacheData).key

		// delete the node pointer in the hash map by key
		delete(c.m, idx)

		// remove the list node
		c.l.Remove(node)
	}
}

func (c *Cache) RemoveMultiple(keys []string) {
	for _, key := range keys {
		c.Remove(key)
	}
}

func (c *Cache) Purge() {
	mp := c.m
	for key := range mp {
		c.Remove(key)
	}
}

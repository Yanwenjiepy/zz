package concurrent_access

import "sync"

type cache struct {
	l sync.RWMutex
	m map[int64]string
}

type ShortIDCache interface {
	Load(shortID int64) (string, bool)
	Store(shortID int64, taskID string)
}

func NewShortIDCache() ShortIDCache {
	sidCache := cache{
		l: sync.RWMutex{},
		m: map[int64]string{},
	}

	return &sidCache
}

func (c *cache) Load(shortID int64) (string, bool) {

	c.l.RLock()
	taskID, ok := c.m[shortID]
	c.l.RUnlock()

	if !ok {
		return "", false
	}

	return taskID, true
}

func (c *cache) Store(shortID int64, taskID string) {
	c.l.Lock()
	c.m[shortID] = taskID
	c.l.Unlock()

}

package concurrent_access

import (
	"sync"
)

const (
	shortIDMax = 9999999
)

// shortID range from 1 to 9999999, include 9999999
type shortID struct {
	l  sync.Mutex
	id int64
}

// ShortIDGenerate generate short id
type GenerateShortID interface {
	Next() int64
}

// Next return next short id
func (sid *shortID) Next() int64 {
	sid.l.Lock()
	defer sid.l.Unlock()

	if sid.id == shortIDMax {
		sid.id = 0
	}

	sid.id += 1

	return sid.id
}

// NewShortIDGenerator return a initialize short id generator
func NewShortIDGenerator() GenerateShortID {

	shortIDGenerator := shortID{
		l:  sync.Mutex{},
		id: 0,
	}

	return &shortIDGenerator
}

package gocache

import (
	"time"

	"github.com/orcaman/concurrent-map"
)

type MapCache struct {
	sc      KVCache
	content cmap.ConcurrentMap
	ticker  *time.Ticker
}

func (mc MapCache) has(k string) bool {
	_, ok := mc.content.Get(k)
	return ok
}

func (mc MapCache) Has(k string) bool {
	if mc.has(k) {
		return true
	} else {
		if mc.sc != nil {
			if mc.sc.Has(k) {
				if value, ok := mc.sc.Get(k); ok {
					mc.set(k, value)
					return true
				}
			}
		}
		return false
	}
}

func (mc MapCache) Get(k string) (interface{}, bool) {
	if mc.Has(k) {
		if tmp, ok := mc.content.Get(k); ok {
			return tmp.(interface{}), true
		} else {
			return nil, false
		}
	}
	return nil, false
}

func (mc MapCache) set(k string, v interface{}) {
	mc.content.Set(k, v)
}

func (mc MapCache) Set(k string, v interface{}) {
	mc.sc.Set(k, v)
	mc.set(k, v)
}

func (mc MapCache) Del(k string) {
	mc.content.Remove(k)
}

func NewMapCache(_sc KVCache) MapCache {
	return MapCache{
		sc:      _sc,
		content: cmap.New(),
	}
}

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

func (mc MapCache) has(k Key) bool {
	_, ok := mc.content.Get(k)
	return ok
}

func (mc MapCache) Has(k Key) bool {
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

func (mc MapCache) Get(k Key) (Value, bool) {
	if mc.Has(k) {
		if tmp, ok := mc.content.Get(k).(Value); ok {
			return tmp, true
		} else {
			return nil, false
		}
	}
	return nil, false
}

func (mc MapCache) set(k Key, v Value) {
	mc.content.Set(k, v)
}

func (mc MapCache) Set(k Key, v Value) {
	mc.sc.Set(k, v)
	mc.set(k, v)
}

func (mc MapCache) Del(k Key) {
	mc.content.Remove(k)
}

func NewMapCache(_sc MapCache) {
	return MapCache{
		sc:      _sc,
		content: cmap.New(),
	}
}

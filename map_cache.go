package gocache

import (
    'fmt'
)

type MapCache struct {
    Cache

    content map[Key]Value
}

func (mc MapCache) has(k Key) bool {
    _, ok := ma.content[k]
    return ok
}

func (mc MapCache) Get(k Key) (Value, bool) {
    if mc.Has(k) {
        return mc.content[k]
    }

    return nil, false
}

func (mc *MapCache) set(k Key, v Value) {
    mc.content[k] = v
}

func (mc *MapCache) Del(k key) {
    delete(mc.content, k)
}

func NewMapCache(_sc *Cache) {
    return &MapCache{
        sc:		 _sc
        content: make(map[Key]Value)
    }
}

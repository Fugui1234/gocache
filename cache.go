package gocache

type string Key
type interface{} Value

type KVCache interface {
    Has(k Key) bool
    Get(k Key) (Value, bool)
    Set(k Key, v Value)
    Del(k Key)

    has(k Key) bool
    set(k Key, v Value)
}

type Cache struct {
    sc *Cache    // secondary cache
}

func (c Cache) Has(k Key) bool {
    if c.has(k) {
        return true
    }

    if c.sc != nil && c.sc.has(k) {
        c.Set(k, c.sc.Get(k))
        return true
    }

    return false
}

func (c *Cache) Set(k Key, v Value) {
    c.set(k, v)
    if c.sc != nil {
        c.sc.set(k, v)
    }
}

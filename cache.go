package gocache

type Key string
type Value interface{}

type KVCache interface {
	Has(k Key) bool
	Get(k Key) (Value, bool)
	Set(k Key, v Value) //单向数据流，不需要此接口
	Del(k Key)

	has(k Key) bool
	set(k Key, v Value)
}

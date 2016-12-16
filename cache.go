package gocache

//type Key string
//type Value interface{}

type KVCache interface {
	Has(k string) bool
	Get(k string) (interface{}, bool)
	Set(k string, v interface{}) //单向数据流，不需要此接口
	Del(k string)

	has(k string) bool
	set(k string, v interface{})
}

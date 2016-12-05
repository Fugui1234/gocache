package gocache

import (
	"testing"
	"fmt"
)

func Test_map_cache(t *testing.T) {
	mc = NewMapCache(nil)

	mc.Set("foo", 0)

	v, ok := mc.Get("foo")
	fmt.Println("v is", v)
}

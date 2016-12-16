package gocache

import (
	"strings"
	//	"time"

	. "../service"
	"github.com/orcaman/concurrent-map"
	"gopkg.in/mgo.v2/bson"
)

type MongoCache struct {
	MapCache
	MC         *MongoC
	fieldNames []string
	table      string
}

func (mc MongoCache) Has(k string) bool {
	if mc.has(k) {
		return true
	} else {
		fields := strings.Split(k, `_|_`)
		query := bson.M{}
		for index, field := range fields {
			query[mc.fieldNames[index]] = field
		}
		results := []bson.M{}
		mc.MC.QueryAll(mc.table, query, &results)

		//只有一个返回结果
		for _, result := range results {
			mc.set(k, result)
			return true
		}
		return false
	}
}

func NewMongoCache(_sc KVCache, _mc *MongoC, names []string, table string) MongoCache {
	return MongoCache{
		MapCache{sc: _sc, content: cmap.New()},
		_mc,
		names,
		table,
	}
}

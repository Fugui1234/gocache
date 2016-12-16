package gocache

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/garyburd/redigo/redis"
	"github.com/orcaman/concurrent-map"
)

type RedisCache struct {
	MapCache
	RC    *redis.Pool
	split string
}

func (rc RedisCache) Has(k string) bool {
	if rc.has(k) {
		return true
	} else {
		key := rc.split + k
		conn := rc.RC.Get()
		defer conn.Close()
		if r, err := redis.String(conn.Do("GET", key)); err == nil {
			res := bson.M{}
			if err := json.Unmarshal([]byte(r), &res); err == nil {
				cconf := bson.M{}
				cconf["res"] = res
				cconf["verifyTime"] = time.Now().Unix()
				cconf["channelId"] = k
				rc.set(k, cconf)
				return true
			}
		}
		return false
	}
}

func NewRedisCache(_sc KVCache, rc *redis.Pool, split string) RedisCache {
	return RedisCache{
		MapCache{sc: _sc, content: cmap.New()},
		rc,
		split,
	}
}

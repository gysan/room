package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gysan/room/config"
)

var goRedis redis.Conn

func GetRedis() redis.Conn {
	var err error
	if goRedis == nil {
		goRedis, err = redis.Dial("tcp", config.RedisBind)
		if err != nil {
			panic(err)
		}
	}
	return goRedis
}

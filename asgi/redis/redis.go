package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisPool is a pool of redis connections
var redisPool *redis.Pool

// CreateRedisPool sets the redis pool to connect to the host.
func CreateRedisPool(host string) {
	redisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", host) },
	}
}

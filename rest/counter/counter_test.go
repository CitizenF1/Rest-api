package counter

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

var redisServer *miniredis.Miniredis

func TestCounterAdd(t *testing.T) {
	setup()
	defer redisServer.Close()

	add := counteradd("10")
	t.Errorf("Counteradd(%s) =>, want %d, %d", "10", 10, add)
}

func setup() {
	redisServer = mockRedis()
	RedisConnector = redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
}

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

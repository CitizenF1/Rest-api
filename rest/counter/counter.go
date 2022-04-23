package counter

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var RedisConnector *redis.Client

func CounterAdd(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	value := counteradd(params["i"])
	fmt.Fprint(w, value)
}

func counteradd(add string) int {
	var value int
	// GET key before add for check value
	var v, _ = RedisConnector.Get("key").Result()
	if v != "" {
		value, _ = strconv.Atoi(v)
		count, _ := strconv.Atoi(add)
		value = value + count
		RedisConnector.Set("key", value, 0).Err()
	} else {
		err := RedisConnector.Set("key", add, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return value
}

func CounterSub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var value int

	// GET key before sub for check value
	var sub, _ = RedisConnector.Get("key").Result()
	if sub != "" {
		value, _ = strconv.Atoi(sub)
		count, _ := strconv.Atoi(params["i"])
		value = value - count
		RedisConnector.Set("key", value, 0).Err()
	} else {
		err := RedisConnector.Set("key", params["i"], 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Fprint(w, value)
}

func CounterVal(w http.ResponseWriter, r *http.Request) {
	// Get value
	val, err := RedisConnector.Get("key").Result()
	if err != nil {
		fmt.Fprint(w, "Redis key is nil")
	}
	fmt.Fprint(w, val)
}

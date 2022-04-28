package counter

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/database/db"
	"strconv"

	"github.com/gorilla/mux"
)

func CounterAdd(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	value := counteradd(params["i"])
	fmt.Fprint(w, value)
}

func counteradd(add string) int {
	var totalSum int
	// GET key before add for check value
	redisConnector := db.RedisConnect()
	var valueKey, err = redisConnector.Get("key").Result()
	if err != nil {
		log.Println("Error Get value from Redis", err)
	}
	// Check value not nil
	if valueKey != "" {
		totalSum, err = strconv.Atoi(valueKey)
		if err != nil {
			fmt.Println(err)
		}
		count, err := strconv.Atoi(add)
		if err != nil {
			fmt.Println(err)
		}
		totalSum = totalSum + count
		err = redisConnector.Set("key", totalSum, 0).Err()
		if err != nil {
			log.Println("Error set value", err)
		}
	} else {
		err := redisConnector.Set("key", add, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	return totalSum
}

func CounterSub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var value int
	redisConnector := db.RedisConnect()
	// GET key before sub for check value
	var sub, _ = redisConnector.Get("key").Result()
	if sub != "" {
		value, _ = strconv.Atoi(sub)
		count, _ := strconv.Atoi(params["i"])
		value = value - count
		redisConnector.Set("key", value, 0).Err()
	} else {
		err := redisConnector.Set("key", params["i"], 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Fprint(w, value)
}

func CounterVal(w http.ResponseWriter, r *http.Request) {
	// Get value
	redisConnector := db.RedisConnect()
	val, err := redisConnector.Get("key").Result()
	if err != nil {
		fmt.Fprint(w, "Redis key is nil")
	}
	fmt.Fprint(w, val)
}

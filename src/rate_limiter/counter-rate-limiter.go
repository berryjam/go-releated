package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

var redisdb *redis.Client

var counterLuaScript = `
	redis.pcall("zadd", KEYS[1], ARGV[1], ARGV[1]); 
	redis.pcall("zremrangebyscore", KEYS[1], 0, ARGV[2]); 
	local count = redis.pcall("zcard", KEYS[1]); 
	redis.pcall("expire", KEYS[1], ARGV[3]); return count`

var evalSha string

func init() {
	initRedisClient()
}

func initRedisClient() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var err error
	evalSha, err = redisdb.ScriptLoad(counterLuaScript).Result()
	if err != nil {
		panic(err)
	}
}

// period单位为秒
func isAllowed(uid string, action string, period, maxCount int) bool {
	key := fmt.Sprintf("%v_%v", uid, action)
	now := time.Now().UnixNano()
	beforeTime := now - int64(period*1000000000)
	res, err := redisdb.EvalSha(evalSha, []string{key}, now, beforeTime, period).Result()
	if err != nil {
		panic(err)
	}
	if res.(int64) > int64(maxCount) {
		return false
	}
	return true
}

func CreateOrder() {
	canCreateOrder := isAllowed("berryjam", "createOrder", 5, 10)
	if canCreateOrder {
		// 处理下单逻辑
		// ...
		fmt.Println("下单成功")
	} else { // 返回请求或者抛出异常、panic
		panic("下单次数超限")
	}
}

func main() {
	for i := 0; i < 100; i++ {
		CreateOrder()
	}
}

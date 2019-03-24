package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

var redisdb *redis.Client

var counterLuaScript = `
	-- 记录行为
	redis.pcall("zadd", KEYS[1], ARGV[1], ARGV[1]); -- value 和 score 都使用纳秒时间戳，即ARGV[1]
	redis.pcall("zremrangebyscore", KEYS[1], 0, ARGV[2]); -- 移除时间窗口之前的行为记录，剩下的都是时间窗口内的
	local count = redis.pcall("zcard", KEYS[1]); -- 获取窗口内的行为数量
	redis.pcall("expire", KEYS[1], ARGV[3]); -- 设置 zset 过期时间，避免冷用户持续占用内存
	return count -- 返回窗口内行为数量`

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

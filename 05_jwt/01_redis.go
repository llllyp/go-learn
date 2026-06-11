package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {

	var ctx = context.Background()

	rdb, err := ConnRedis("localhost:6379", "123456", 0)
	if err != nil {
		panic(err)
	}

	defer rdb.Close()
	// Sting  set/get
	var redisKeyStr = "user:1:name:string"
	err = rdb.Set(ctx, redisKeyStr, "java_dev", 0).Err()
	if err != nil {
		panic(err)
	}

	value, err := rdb.Get(ctx, redisKeyStr).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("string value: " + value)

	// Hash
	redisKeyHash := "user:1:name:hash"
	rdb.HSet(ctx, redisKeyHash, map[string]interface{}{
		"id":   1,
		"name": "java_dev",
		"age":  28,
	})

	name, _ := rdb.HGet(ctx, redisKeyHash, "name").Result()
	fmt.Println("hash value, name: " + name)

	// 批量操作
	pipe := rdb.Pipeline()
	pipe.Set(ctx, "k1", "v1", 0)
	pipe.Set(ctx, "k2", "v2", 0)
	pipe.Set(ctx, "k3", "v3", 0)

	_, err = pipe.Exec(ctx)

	// Lua 原子操作
	luaScript := `
	local key = KEYS[1]
	local limit = tonumber(ARGV[1])
	local current = tonumber(redis.call("GET", key) or "0")
	
	if current + 1 > limit then
	    return 0
	end
	
	redis.call("INCR", key)
	redis.call("EXPIRE", key, ARGV[2])
	return 1
	`
	res, err := rdb.Eval(ctx, luaScript, []string{"rate:limit"}, 10, 60).Result()
	fmt.Println(res)
	
}

func ConnRedis(addr, pass string, dbNum int8) (*redis.Client, error) {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       int(dbNum),
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	// https://app.redislabs.com/#/subscriptions/subscription/2241897/bdb
	options, err := redis.ParseURL("redis://default:Xhuym3KS1o1ydaAMdhwQ4EEhvPxSwbtd@redis-18260.c274.us-east-1-3.ec2.cloud.redislabs.com:18260")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(options)

	defer rdb.Close()

	ctx := context.Background()

	err = rdb.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	key1Res, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	println(key1Res)

	err = rdb.MSet(ctx, "key2", 200, "key3", 300).Err()
	if err != nil {
		panic(err)
	}

	key3Res, err := Get(ctx, rdb, "key3").Result()
	if err == redis.Nil {
		println("field not exit")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(key3Res)

	exist, err := rdb.Exists(ctx, "key4").Result()
	if err != nil {
		panic(err)
	}
	if exist == 0 {
		fmt.Println("key4 is not exist")
		err = rdb.Set(ctx, "key4", 4444, 0).Err()
		if err != nil {
			panic(err)
		}
	}

	deleted, err := rdb.Del(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	if deleted == 1 {
		fmt.Println("Key deleted")
	} else {
		fmt.Println("Key does not exist")
	}

	// list type
	err = rdb.LPush(ctx, "List", "value1", "value2").Err()
	if err != nil {
		panic(err)
	}

	// set type
	err = rdb.SAdd(ctx, "set", "value1", "value2").Err()
	if err != nil {
		panic(err)
	}
}

func Get(ctx context.Context, rdb *redis.Client, key string) *redis.StringCmd {
	cmd := redis.NewStringCmd(ctx, "get", key)
	err := rdb.Process(ctx, cmd)
	if err != nil {
		panic(err)
	}
	return cmd
}

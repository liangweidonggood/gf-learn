package main

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/util/gconv"
)

var (
	config = gredis.Config{
		Host: "192.168.230.129",
		Port: 6379,
		Pass: "ZZkde@#3d99",
		Db:   1,
	}
)

func main() {
	group := "test"
	gredis.SetConfig(&config, group)

	redis := gredis.Instance(group)
	defer redis.Close()

	_, err := redis.Do("SET", "k", "v")
	if err != nil {
		panic(err)
	}

	r, err := redis.Do("GET", "k")
	if err != nil {
		panic(err)
	}
	fmt.Println(gconv.String(r))
}

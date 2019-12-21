package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
)

var Pool *redis.Pool

func init() {
	redisHost := ":6379"
	Pool = newPool(redisHost)
	closePool()
}

func newPool(serve string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   8, //最大空闲链接数
		MaxActive: 0, // 表示和数据库的最大链接数， 0 表示没有限制 IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个 ip 的 redis
			return redis.Dial("tcp", serve)
		},
	}
}

func closePool() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func main() {
	// go 连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	fmt.Println("connect success")

	_, err = conn.Do("Set", "name", "leighj")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	// 返回的r是interface{}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
	}
	// r = r.(string)  panic: interface conversion: interface {} is []uint8, not string
	fmt.Println(r)
	//设置过期时间
	_, err = conn.Do("expire", "name", 10)

	_, err = conn.Do("HSet", "user01", "name", "leighj")
	r, err = redis.String(conn.Do("HGet", "user01", "name"))
	_, err = conn.Do("HSet", "user01", "age", 18)
	i, err := redis.Int(conn.Do("HGet", "user01", "age"))
	//list
	_, err = conn.Do("lpush", "heroList", "no1:宋江", 30, "no2:卢俊义", 28)
	r, err = redis.String(conn.Do("rpop", "heroList"))

	fmt.Println(r)
	fmt.Println(i)
}

package db

import (
	"log"
	"github.com/gomodule/redigo/redis"
	"fmt"
)

var (
	Conn redis.Conn
)

func init() {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal("error connect to redis!")
	}
	Conn = c
	if _, err := Conn.Do("SELECT", 2); err != nil {
		Conn.Close()
	}
}

func Get(key string) ([]string, error) {
    var data []string
	data, err := redis.Strings(Conn.Do("SMEMBERS", key))
    if err != nil {
        return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
    return data, err
}
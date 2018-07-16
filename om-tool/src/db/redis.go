package db

import (
	d "om-tool/src/models"
	"strings"
	"log"
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func Ping() error {
    conn := Pool.Get()
    defer conn.Close()

    _, err := redis.String(conn.Do("PING"))
    if err != nil {
        return fmt.Errorf("cannot 'PING' db: %v", err)
    }
    return nil
}

func Get(key string) ([]byte, error) {

    conn := Pool.Get()
    defer conn.Close()

    var data []byte
    data, err := redis.Bytes(conn.Do("GET", key))
    if err != nil {
        return data, fmt.Errorf("error getting key %s: %v", key, err)
    }
    return data, err
}

func Set(key string, value []byte) error {
    conn := Pool.Get()
    defer conn.Close()
    _, err := conn.Do("SET", key, value)
    if err != nil {
        v := string(value)
        return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
    }
    return err
}

func MSet(kvs []d.KeyValue) {
    conn := Pool.Get()
    defer conn.Close()
    conn.Send("MULTI")
    for _, kv := range kvs {
        if strings.Replace(kv.Key, " ", "", -1) != "" && len(kv.Value) > 0 {
            if err := conn.Send("SADD", kv.Key, kv.Value); err != nil {
                log.Println(err)
            }
        } else {
            log.Println(kv.Key)
        }
    }
    conn.Do("EXEC")
}

func MGet(keys []string) map[string][]byte {
    conn := Pool.Get()
    defer conn.Close()
    conn.Send("MULTI")
    for _, key := range keys {
        if strings.Replace(key, " ", "", -1) != "" {
            if err := conn.Send("GET", key); err != nil {
                log.Fatal(err)
            }
        }
    }
    results, _ := redis.Values(conn.Do("EXEC"))
    kvs := make(map[string][]byte)
    for i, result := range results {
        r := []byte{}
        if result != nil {
            r, _ = redis.Bytes(result, nil)
            kvs[keys[i]] = r
        }
    }
    return kvs
}

func transStringToInterface(keys []string) []interface{} {
    new := make([]interface{}, len(keys))
    for i, v := range keys {
        new[i] = v
    }
    return new
}

func transIdToInterface(keys []d.IdentifiersObj) []interface{} {
    new := make([]interface{}, len(keys))
    for i, v := range keys {
        new[i] = v
    }
    return new
}

func Exists(key string) (bool, error) {
    conn := Pool.Get()
    defer conn.Close()

    ok, err := redis.Bool(conn.Do("EXISTS", key))
    if err != nil {
        return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
    }
    return ok, err
}

func Delete(key string) error {

    conn := Pool.Get()
    defer conn.Close()

    _, err := conn.Do("DEL", key)
    return err
}

func GetKeys(pattern string) ([]string, error) {

    conn := Pool.Get()
    defer conn.Close()

    iter := 0
    keys := []string{}
    for {
        arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
        if err != nil {
            return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
        }

        iter, _ = redis.Int(arr[0], nil)
        k, _ := redis.Strings(arr[1], nil)
        keys = append(keys, k...)

        if iter == 0 {
            break
        }
    }

    return keys, nil
}

func Incr(counterKey string) (int, error) {

    conn := Pool.Get()
    defer conn.Close()

    return redis.Int(conn.Do("INCR", counterKey))
}
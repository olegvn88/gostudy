package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/olegvn88/gostudy/app/main/app/database/properties"
	"time"
)

//redis - это такой memcached на стероидах ))

var (
	c    redis.Conn
	prop = properties.JsonProp()
)

//
//func getRecord(mKey string) string {
//	println("get", mKey)
//	data, err := c.Do("GET", mKey)
//	item, err := redis.String(data, err)
//	//если записи нет, то для этого есть специальная ошибка, ее нужно обрабатывать отдельно
//	 if err == redis.ErrNil {
//	 	fmt.Println("Record is not found in redis (return value is nil)")
//	 	return ""
//	 } else if err != nil {
//	 	PanicOnError(err)
//	 }
//	return item
//}

func getRecord(mkey string) (string, error) {
	println("get", mkey)
	// получает запись, https://redis.io/commands/get
	data, err := c.Do("GET", mkey)
	item, err := redis.String(data, err)
	// если записи нету, то для этого есть специальная ошибка, её надо обрабатывать отдеьно, это почти штатная ситуация, а не что-то страшное
	if err == redis.ErrNil {
		fmt.Println("Record not found in redis (return value is nil)")
		return "", redis.ErrNil
	} else if err != nil {
		return "", err
		PanicOnError(err)
	}
	return item, nil
}

func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var err error
	c, err = redis.DialURL("redis://:" + prop.RedisPassword + "@" + prop.RedisHost)
	//c, err := redis.DialURL("redis://" + prop.RedisHost)
	PanicOnError(err)
	defer c.Close()

	mKey := "record_21"
	item, err := getRecord(mKey)
	fmt.Printf("first get %+v\n", item)

	//добавляет запись
	ttl := 5
	result, err := redis.String(c.Do("SET", mKey, 1, "EX", ttl)) //EX = expiration
	if result != "OK" {
		panic("result is not ok " + result)
	}

	time.Sleep(time.Microsecond)

	item, err = getRecord(mKey)
	fmt.Printf("second get %+v\n", item)

	n, _ := redis.Int(c.Do("DECRBY", mKey, 1))
	fmt.Println("DECRY by 1 ", mKey, "is", n)

	//если записи не было - редис создаст
	n, err = redis.Int(c.Do("INCR", mKey+"_not_exist"))
	fmt.Println("INCR (default by 1) ", mKey+"_not_exist", "is", n)

	PanicOnError(err)

	keys := []interface{}{mKey, mKey + "_not_exist", "sure_not_exist"}

	reply, err := redis.Strings(c.Do("MGET", keys...))
	PanicOnError(err)
	fmt.Println(reply)

}

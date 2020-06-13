package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

type AppCache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) AppCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (cache *redisCache) getClient() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Get(key string) interface{} {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var res interface{}

	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return res
}

func (cache *redisCache) Set(key string, value interface{}) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	client.Set(key, json, cache.expires*time.Second)
}

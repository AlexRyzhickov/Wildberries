package main

import (
	"context"
	"errors"
)

const (
	redis     = "redis"
	memcached = "memcached"
)

type CacheI interface {
	Store(ctx context.Context, key string, value interface{}) error
	Load(ctx context.Context, key string) (interface{}, bool)
	Delete(ctx context.Context, key string) error
}

type RedisCache struct {
}

func newRedisCache() CacheI {
	return &RedisCache{}
}

func (c *RedisCache) Store(ctx context.Context, key string, value interface{}) error {
	return nil
}

func (c *RedisCache) Load(ctx context.Context, key string) (interface{}, bool) {
	return nil, true
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return nil
}

type MemcachedCache struct {
}

func newMemcachedCache() CacheI {
	return &MemcachedCache{}
}

func (c *MemcachedCache) Store(ctx context.Context, key string, value interface{}) error {
	return nil
}

func (c *MemcachedCache) Load(ctx context.Context, key string) (interface{}, bool) {
	return nil, true
}

func (c *MemcachedCache) Delete(ctx context.Context, key string) error {
	return nil
}

func initCache(cacheType string) (CacheI, error) {
	switch cacheType {
	case redis:
		return newRedisCache(), nil
	case memcached:
		return newMemcachedCache(), nil
	}
	return nil, errors.New("invalid cache type")
}

func main() {
	redisCache, _ := initCache(redis)
	memCached, _ := initCache(memcached)
	_, _ = redisCache, memCached
}

/*
Фабричный метод — это порождающий паттерн проектирования, который решает проблему создания различных продуктов, без указания конкретных классов продуктов.
+ Позволяет создавать объекты без привязки к класссам
*/

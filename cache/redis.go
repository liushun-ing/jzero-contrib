package cache

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/jsonx"
	zerocache "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
)

type redisNode struct {
	rds  *redis.Redis
	node zerocache.Cache
}

func (c redisNode) SetNoExpireCtx(ctx context.Context, key string, val any) error {
	data, err := jsonx.Marshal(val)
	if err != nil {
		return err
	}
	return c.rds.SetCtx(ctx, key, string(data))
}

func (c redisNode) Del(keys ...string) error {
	return c.node.Del(keys...)
}

func (c redisNode) DelCtx(ctx context.Context, keys ...string) error {
	return c.node.DelCtx(ctx, keys...)
}

func (c redisNode) Get(key string, val any) error {
	return c.node.Get(key, val)
}

func (c redisNode) GetCtx(ctx context.Context, key string, val any) error {
	return c.node.GetCtx(ctx, key, val)
}

func (c redisNode) IsNotFound(err error) bool {
	return c.node.IsNotFound(err)
}

func (c redisNode) Set(key string, val any) error {
	return c.node.SetCtx(context.Background(), key, val)
}

func (c redisNode) SetCtx(ctx context.Context, key string, val any) error {
	return c.node.SetCtx(ctx, key, val)
}

func (c redisNode) SetWithExpire(key string, val any, expire time.Duration) error {
	return c.node.SetWithExpireCtx(context.Background(), key, val, expire)
}

func (c redisNode) SetWithExpireCtx(ctx context.Context, key string, val any, expire time.Duration) error {
	return c.node.SetWithExpireCtx(ctx, key, val, expire)
}

func (c redisNode) Take(val any, key string, query func(val any) error) error {
	return c.node.Take(val, key, query)
}

func (c redisNode) TakeCtx(ctx context.Context, val any, key string, query func(val any) error) error {
	return c.node.TakeCtx(ctx, val, key, query)
}

func (c redisNode) TakeWithExpire(val any, key string, query func(val any, expire time.Duration) error) error {
	return c.node.TakeWithExpire(val, key, query)
}

func (c redisNode) TakeWithExpireCtx(ctx context.Context, val any, key string, query func(val any, expire time.Duration) error) error {
	return c.node.TakeWithExpireCtx(ctx, val, key, query)
}

func NewRedisNode(rds *redis.Redis, errNotFound error, opts ...zerocache.Option) Cache {
	singleFlights := syncx.NewSingleFlight()
	stats := zerocache.NewStat("redis-cache")
	node := zerocache.NewNode(rds, singleFlights, stats, errNotFound, opts...)
	return &redisNode{
		rds:  rds,
		node: node,
	}
}

package data

import (
	"context"
	"gin-jj/app"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     app.C.Redis.Host + ":" + app.C.Redis.Port,
		Password: app.C.Redis.Password,
		DB:       app.C.Redis.Database,
	})
	var ctx = context.Background()
	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		app.Log.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		app.Log.Info("redis connect ping response:", zap.String("pong", pong))
		app.Redis = rdb
	}
}

func SetRedis(key string, value string, t int64) bool {
	var ctx = context.Background()
	expire := time.Duration(t) * time.Second
	if err := rdb.Set(ctx, key, value, expire).Err(); err != nil {
		return false
	}

	return true
}

func GetRedis(key string) string {
	var ctx = context.Background()

	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}

	return res
}

func DelRedis(key string) bool {
	var ctx = context.Background()
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return false
	}
	return true
}

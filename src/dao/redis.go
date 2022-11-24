/**
* @Description：
* @Author: cdx
* @Date: 2022/11/24 2:20 下午
 */

package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"short_url/conf"
	"time"
)

var RedisClient RedisSrv

type RedisSrv struct {
	client *redis.Client
}

func newRedisClient() {

	conf := conf.GetConfig().RedisConf
	fmt.Println("conf = ", conf)
	RedisClient.client = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.PassWd,
		DB:       conf.Db,
	})

}

func (r *RedisSrv) Set(ctx context.Context) (err error) {

	resp, err := r.client.Set(ctx, "111", "2222", time.Duration(10*time.Second)).Result()
	fmt.Println("set resp=", resp, ",err=", err)

	val2, err := r.client.Get(ctx, "111").Result()
	fmt.Println("get resp=", val2, ",err=", err)

	return
}

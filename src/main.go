/**
* @Description：
* @Author: cdx
* @Date: 2022/11/24 12:09 下午
 */

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"short_url/dao"
	"short_url/middleware/logger"
	"short_url/server"
)

func main() {

	dao.Init()

	dao.RedisClient.Set(context.Background())

	router()
}

func router() {

	r := gin.New()

	r.Use(logger.Logger())
	r.Use(gin.Recovery())

	r.GET("/urlEncode/:url", server.HandleUrlEncode)

	r.Run(":3333")
}

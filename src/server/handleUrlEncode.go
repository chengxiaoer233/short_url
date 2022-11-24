/**
* @Description：
* @Author: cdx
* @Date: 2022/11/24 5:47 下午
 */

package server

import (
	"github.com/gin-gonic/gin"
	"short_url/middleware/logger"

	"net/http"
	"short_url/model"
)

// 长连接生成对应的短连接
func HandleUrlEncode(c *gin.Context) {

	var respData model.RespStruct

	// input
	param := c.Param("url")
	if param == "" {
		c.JSON(http.StatusOK, "param url empty")
		return
	}

	// handle
	shortUrl, err := encodeUrl(param)
	if err != nil {

		logger.Log().Errorf("encodeUrl error,err=%s", err)
		c.JSON(http.StatusOK, model.RespStruct{Code: 1000, Msg: err.Error()})
		return
	}

	respData.Code = 0
	respData.Msg = "success"
	respData.OriginUrl = param
	respData.ShortUrl = shortUrl
	c.JSON(http.StatusOK, respData)

	return
}

func encodeUrl(originUrl string) (shortUrl string, err error) {

	return "", nil
}

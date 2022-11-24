/**
* @Description：
* @Author: cdx
* @Date: 2022/11/24 6:01 下午
 */

package model

type RespStruct struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	OriginUrl string `json:"originUrl"`
	ShortUrl  string `json:"shortUrl"`
}

package main

import (
	"encoding/xml"
	"fmt"
	"github.com/BruceMaa/Panda/wechat/mp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化配置微信
var wechatMp = &mp.WechatMp{}

func init() {
	wechatMpConfig := &ChatConfig.WechatConfig.WechatMpConfig
	wechatMp = mp.New(wechatMpConfig)
}

// 设置微信公众号服务
func wechatMpServer(router *gin.Engine) {
	wechatApis := router.Group("/wechatmp")
	{
		wechatApis.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, wechatMp.AuthWechatServer(c.Request))
		})

		wechatApis.POST("", func(c *gin.Context) {
			c.String(http.StatusOK, wechatMp.CallBackFunc(c.Request))
		})
	}

	//wechatMp.SetTextHandlerFunc(TextMessageEncryptHandler)
	wechatMp.SetTextHandlerFunc(TextMessageHandler)
}

func TextMessageHandler(request *mp.TextMessage) string {
	fmt.Printf("text message: %+v\n", request)
	textResp, _ := mp.NewMsgTextResponseString(request.FromUserName, request.ToUserName, request.Content)
	fmt.Printf("textResp: %s\n", textResp)
	return textResp
}

func TextMessageEncryptHandler(request *mp.TextMessage) string {
	fmt.Printf("text message: %+v\n", request)
	textRespByte, _ := mp.NewMsgTextResponseByte(request.FromUserName, request.ToUserName, request.Content)
	encryptRespByte, _ := wechatMp.AESEncryptMessage(textRespByte)
	fmt.Printf("encryptRespByte: %+v\n", encryptRespByte)
	encryptResp, _ := xml.Marshal(encryptRespByte)
	fmt.Printf("encryptResp: %s\n", encryptResp)
	return string(encryptResp)
}

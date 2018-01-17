package main

import (
	"fmt"
	"github.com/BruceMaa/gokit/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var timeout = 5 * time.Second

func main() {

	app := &http.Server{
		Addr:         fmt.Sprintf(":%d", ChatConfig.Port),
		Handler:      initRouter(),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	Logger.Debug("当前运行环境：", gin.Mode())

	tool.StartAndQuitHTTPGraceful(app, timeout)
}

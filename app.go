package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	// 配置日志，日志轮询切分，可以使用Linux系统中的logrotate
	//logFile, _ := os.Create("gochat.log")
	//errorLogFile, _ := os.Create("gochat_error.log")
	//DebugHandle = io.MultiWriter(logFile)
	//TraceHandle = io.MultiWriter(logFile)
	//InfoHandle = io.MultiWriter(logFile)
	//WarnHandle = io.MultiWriter(logFile)
	//ErrorHandle = io.MultiWriter(logFile)
	//
	//gin.DefaultWriter = io.MultiWriter(logFile)
	//gin.DefaultErrorWriter = io.MultiWriter(errorLogFile)
}

func main() {

	app := &http.Server{
		Addr:         fmt.Sprintf(":%d", ChatConfig.Port),
		Handler:      initRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	Logger.Debug("当前运行环境：", gin.Mode())

	app.ListenAndServe()

}

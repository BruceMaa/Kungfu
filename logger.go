package main

import (
	"github.com/BruceMaa/gokit/tool"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var Logger = &tool.Logger{}

func init() {
	// 配置日志，日志轮询切分，可以使用Linux系统中的logrotate
	logFile, _ := os.Create("kungfu.log")
	errorLogFile, _ := os.Create("kungfu_error.log")
	tool.DebugHandle = io.MultiWriter(logFile)
	tool.TraceHandle = io.MultiWriter(logFile)
	tool.InfoHandle = io.MultiWriter(logFile)
	tool.WarnHandle = io.MultiWriter(logFile)
	tool.ErrorHandle = io.MultiWriter(logFile)

	gin.DefaultWriter = io.MultiWriter(logFile)
	gin.DefaultErrorWriter = io.MultiWriter(errorLogFile)
}

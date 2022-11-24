package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"short_url/conf"
	"time"
)

var logger *logrus.Logger

func Log() *logrus.Logger {
	return logger
}

func Logger() gin.HandlerFunc {

	//linkName := "latest_log.log"
	filePath := conf.GetConfig().LogStruct.LogPath      // 日志保存位置
	maxAge := conf.GetConfig().LogStruct.MaxAge         // 日志保存最大时间
	rotateTime := conf.GetConfig().LogStruct.RotateTime // 日志切割时间
	level := conf.GetConfig().LogStruct.Level           // 日志级别

	scr, err := os.OpenFile(filePath+"/log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(fmt.Sprintf("logFile open err=%s", err.Error()))
	}

	logger = logrus.New()
	logger.Out = scr
	logger.SetLevel(logrus.Level(level))

	// 日志属性设置
	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(time.Duration(maxAge)*time.Hour),
		retalog.WithRotationTime(time.Duration(rotateTime)*time.Hour),
		//retalog.WithLinkName(linkName),
	)

	// 写入的日志级别
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	// 需要写入的日志字段
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime).Milliseconds()
		spendTime := fmt.Sprintf("%d ms", stopTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		// 根据对应的code码来生成对应的日志
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

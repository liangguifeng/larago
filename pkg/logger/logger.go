package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"log"
	"os"
	"strings"
)

// LogError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Debug(format string) {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%",
		},
	}

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	logger.Debug(format)
}

func DebugPrint(a ...interface{}) {
	fmt.Printf("")
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%",
		},
	}

	logger.Debug(a...)
}

func Info(format string) {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%",
		},
	}

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	logger.Info(format)
}

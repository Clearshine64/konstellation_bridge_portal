package logger

import (
	"io"
	"os"

	echoLog "github.com/labstack/gommon/log"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
)

const LogFileName = "log.txt"

var (
	applog      *AppLogger
	LogFile     *os.File
	MultyWriter io.Writer
)

type AppLogger struct {
	*log.MyLogger
}

// Log singleton instance
func Log() *AppLogger {
	return &AppLogger{log.Logger()}
}

// it's just logger interface idea
// ----------------------------------------------------------------------
// type Logger interface {
// 	// should be used for info log
// 	Info(format string, args ...interface{})
// 	// should be used for warning log
// 	Warn(format string, args ...interface{})
// 	// should be used for error log
// 	Err(format string, args ...interface{})
// 	// should be used for debug
// 	Debug(args ...interface{})
// 	// should be used for panic
// 	Panic(msg string)
// 	// should be used for info log with params
// 	InfoP(format string, paramsLog().Fields, args ...interface{})
//  // should be used for warn log with params
// 	WarnP(format string, paramsLog().Fields, args ...interface{})
//	// should be used for error log with params
// 	ErrP(format string, paramsLog().Fields, args ...interface{})
// }
// ----------------------------------------------------------------------

// Info log
func Info(format string, args ...interface{}) {
	Log().SetLevel(echoLog.INFO)
	Log().Infof(format, args...)
}

// Warn log
func Warn(format string, args ...interface{}) {
	Log().SetLevel(echoLog.WARN)
	Log().Infof(format, args...)
}

// Err log
func Err(format string, args ...interface{}) {
	Log().SetLevel(echoLog.ERROR)
	Log().Errorf(format, args...)
}

// Debug log
func Debug(args ...interface{}) {
	Log().SetLevel(echoLog.DEBUG)
	Log().Debug(args...)
}

// Panic start
func Panic(i ...interface{}) {
	Log().SetLevel(echoLog.ERROR)
	Log().Panic(i...)
}

// InfoP log
func InfoP(format string, params logrus.Fields, args ...interface{}) {
	Log().SetLevel(echoLog.INFO)
	Log().WithFields(params).Infof(format, args...)
}

// WarnP log
func WarnP(format string, params logrus.Fields, args ...interface{}) {
	Log().SetLevel(echoLog.WARN)
	Log().WithFields(params).Warnf(format, args...)
}

// ErrP log
func ErrP(format string, params logrus.Fields, args ...interface{}) {
	Log().SetLevel(echoLog.ERROR)
	Log().WithFields(params).Errorf(format, args...)
}

func SetLogFile(fileName string) error {
	logFileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	LogFile = logFileObj
	MultyWriter = io.MultiWriter(os.Stdout, LogFile)
	return nil
}

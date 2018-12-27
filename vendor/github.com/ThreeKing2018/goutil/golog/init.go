package golog

import (
	"github.com/ThreeKing2018/goutil/golog/conf"
	"github.com/ThreeKing2018/goutil/golog/plugins/zaplog"
)

//默认
var l ILog = zaplog.New()

type backend uint8

const (
	ZAPLOG backend = iota
	DlOG
)

//设置
func SetLogger(b backend, opts ...conf.Option) {
	switch b {
	case ZAPLOG:
		l = zaplog.New(opts...)
	case DlOG:

	}
}

//目前只有zap生效
func SetLogLevel(level conf.Level) {
	l.SetLogLevel(level)
}

//目前只有zap生效
func Sync() {
	l.Sync()
}

//普通日志
func Debug(args ...interface{}) {
	l.Debug(args...)
}
func Info(args ...interface{}) {
	l.Info(args...)
}
func Warn(args ...interface{}) {
	l.Warn(args...)
}
func Error(args ...interface{}) {
	l.Error(args...)
}
func Panic(args ...interface{}) {
	l.Panic(args...)
}
func Fatal(args ...interface{}) {
	l.Fatal(args...)
}

//需要格式化日志
func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

//key value
func Debugw(msg string, keysAndValues ...interface{}) {
	l.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	l.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	l.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	l.Errorw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	l.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	l.Fatalw(msg, keysAndValues...)
}

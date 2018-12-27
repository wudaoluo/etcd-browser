package golog

import (
	"github.com/ThreeKing2018/goutil/golog/conf"
)

//使用string是为了减少使用Spintf
type ILog interface {
	//普通日志
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Fatal(...interface{})

	//需要格式化日志
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})

	//key value
	Debugw(string, ...interface{})
	Infow(string, ...interface{})
	Warnw(string, ...interface{})
	Errorw(string, ...interface{})
	Panicw(string, ...interface{})
	Fatalw(string, ...interface{})

	Sync()
	SetLogLevel(conf.Level)
}

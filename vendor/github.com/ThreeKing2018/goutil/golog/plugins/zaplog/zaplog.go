package zaplog

import (
	"github.com/ThreeKing2018/goutil/golog/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Log struct {
	logger *zap.SugaredLogger
	atom   zap.AtomicLevel
}

func parseLevel(level conf.Level) zapcore.Level {
	switch level {
	case conf.DebugLevel:
		return zapcore.DebugLevel
	case conf.InfoLevel:
		return zapcore.InfoLevel
	case conf.WarnLevel:
		return zapcore.WarnLevel
	case conf.ErrorLevel:
		return zapcore.ErrorLevel
	case conf.PanicLevel:
		return zapcore.PanicLevel
	case conf.FatalLevel:
		return zapcore.FatalLevel
	}

	return zapcore.DebugLevel
}

var encoderConfig = zapcore.EncoderConfig{
	// Keys can be anything except the empty string.
	TimeKey:        "T",
	LevelKey:       "L",
	NameKey:        "N",
	CallerKey:      "C",
	MessageKey:     "M",
	StacktraceKey:  "S",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func New(opts ...conf.Option) *Log {
	o := &conf.Options{
		Filename:   conf.Filename,
		LogLevel:   conf.LogLevel,
		MaxSize:    conf.MaxSize,
		MaxAge:     conf.MaxAge,
		Stacktrace: conf.Stacktrace,
		IsStdOut:   conf.IsStdOut,
		//ProjectName: conf.ProjectName,
		LogType: conf.LogNormalType,
	}

	for _, opt := range opts {
		opt(o)
	}

	var writers = []zapcore.WriteSyncer{}
	osfileout := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize, // megabytes
		MaxBackups: 3,
		MaxAge:     o.MaxAge, // days
		LocalTime:  true,
	})
	if o.IsStdOut {
		writers = append(writers, os.Stdout)
	}

	writers = append(writers, osfileout)
	w := zapcore.NewMultiWriteSyncer(writers...)

	atom := zap.NewAtomicLevel()
	atom.SetLevel(parseLevel(o.LogLevel)) //改变日志级别

	var enc zapcore.Encoder
	if o.LogType == conf.LogNormalType {
		enc = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		enc = zapcore.NewJSONEncoder(encoderConfig)
	}
	core := zapcore.NewCore(
		//这里控制json 或者不是json 类型
		enc,
		w,
		atom,
	)

	logger := zap.New(
		core,
		zap.AddStacktrace(parseLevel(o.Stacktrace)),
		zap.AddCaller(),
		zap.AddCallerSkip(2))

	if o.ProjectName != "" {
		logger = logger.With(zap.String(conf.ProjectKey, o.ProjectName))
	}
	loggerSugar := logger.Sugar()
	return &Log{logger: loggerSugar, atom: atom}

}

func (l *Log) Sync() {
	l.Sync()
}

func (l *Log) SetLogLevel(level conf.Level) {
	l.atom.SetLevel(parseLevel(level))
}

func (l *Log) Debug(fields ...interface{}) {
	l.logger.Debug(fields)
}
func (l *Log) Info(fields ...interface{}) {
	l.logger.Info(fields)
}
func (l *Log) Warn(fields ...interface{}) {
	l.logger.Warn(fields)
}
func (l *Log) Error(fields ...interface{}) {
	l.logger.Error(fields)
}
func (l *Log) Panic(fields ...interface{}) {
	l.logger.Panic(fields)
}
func (l *Log) Fatal(fields ...interface{}) {
	l.logger.Fatal(fields)
}

func (l *Log) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}
func (l *Log) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}
func (l *Log) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}
func (l *Log) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}
func (l *Log) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}
func (l *Log) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *Log) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}
func (l *Log) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}
func (l *Log) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}
func (l *Log) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}
func (l *Log) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}
func (l *Log) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

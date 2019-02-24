package serverRoom

import (
	"fmt"
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/ThreeKing2018/goutil/golog/conf"
	"os"
	"path"
)



func Init() {
	//flag.Parse()

	//打印版本并退出
	if Arg.Getver() {
		fmt.Println(Version)
		os.Exit(0)

	}
	cnf := GetConfigInstance()

	logType := conf.LogNormalType
	if cnf.GetBool("log_type_json") {
		logType = conf.LogJsontype
	}


	logLevel := conf.WarnLevel
	if cnf.GetBool("debug") {
		logLevel = conf.DebugLevel
	}

	golog.SetLogger(
		golog.ZAPLOG,
		conf.WithLogType(logType),
		conf.WithLogLevel(logLevel),
		conf.WithFilename(path.Join(cnf.GetString("log_dir"),
			cnf.GetString("service_name")+".log")),
	)
}


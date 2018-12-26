package serverRoom

import (
	"flag"
	"fmt"
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/ThreeKing2018/goutil/golog/conf"
	"github.com/wudaoluo/etcd-browser/internal"
	"path"
)



func Init() {
	flag.Parse()

	//打印版本并退出
	if Arg.Getver() {
		fmt.Println(Version)
	}

	golog.SetLogger(
		golog.ZAPLOG,
		conf.WithLogType(conf.LogJsontype),
		conf.WithLogLevel(conf.DebugLevel),
		conf.WithFilename(path.Join(Arg.logdir,internal.SERVICE_NAME+".log")),
	)
}


package serverRoom

import (
	"flag"
)

const (
	VERSION = "0.1"
	DEFAULT_CONFIG_NAME = "config.json"
)

type argStruct struct {
	version bool
	configfile string
	//debug bool
	//logdir string
}

var Arg = new(argStruct)

func init() {
	flag.BoolVar(&Arg.version,"version",false,"打印版本号")
	//flag.BoolVar(&Arg.debug,"debug",true,"open debug default false")
	flag.StringVar(&Arg.configfile,"c",DEFAULT_CONFIG_NAME,"specify config file")
	//flag.StringVar(&Arg.logdir,"logdir",internal.LOG_DIR,"log dir")
}




func (a *argStruct) Getver() bool{
	return a.version
}

func (a *argStruct) GetConfigFile() string {
	return a.configfile
}

//func (a *argStruct) GetDebug() bool {
//	return a.debug
//}

//func (a *argStruct) GetLogDir() string {
//	return a.logdir
//}
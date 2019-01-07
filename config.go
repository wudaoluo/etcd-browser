package serverRoom

import (
	"github.com/ThreeKing2018/goutil/config"
	"strings"
	"sync"
)

//配置文件操作

type singleton config.Viperable


var v singleton
var once sync.Once


func GetConfigInstance() singleton {
	once.Do(load)
	return v
}

//配置文件初始化
func load() {
	v = config.New()
	var cnfType = "json"
	a := strings.Split(Arg.configfile,".")
	if len(a) == 2 {
		cnfType = a[1]
	}
	v.SetConfig(Arg.configfile, cnfType, "/etc", "/home", ".")

	err := v.ReadConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()



}

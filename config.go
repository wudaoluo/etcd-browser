package serverRoom

import (
	"github.com/ThreeKing2018/goutil/config"
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
	v.SetConfig(Arg.configfile, "json", "/etc", "/home", ".")

	err := v.ReadConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()



}

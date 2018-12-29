package serverRoom

import (
	"fmt"
	"github.com/ThreeKing2018/goutil/config"
	"time"
)

//配置文件操作

func init() {
	conf := config.New()
	conf.SetConfig("config.json", "json", "/etc", "/home", ".")
	err := conf.ReadConfig()
	if err != nil {
		panic(err)
	}
	conf.WatchConfig()
	go func() {
		for {
			fmt.Println(conf.GetString("debug"))
			time.Sleep(1 * time.Second)
		}
	}()
}

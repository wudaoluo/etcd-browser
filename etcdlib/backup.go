package etcdlib

import (
	"context"
	"github.com/ThreeKing2018/goutil/config"
	"github.com/ThreeKing2018/goutil/golog"
	"go.etcd.io/etcd/clientv3"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const shortDate = "2006_01_02"

type singleton config.Operater

var v singleton
var once sync.Once

func getOperating() singleton {
	once.Do(load)
	return v
}

//配置文件初始化
func load() {
	var err error
	v, err = config.NewOperater()
	if err != nil {
		golog.Error("初始化备份文件操作失败", "err", err)
		panic(err)
	}
}

func (c *client) Backup(fpath string) {
	if !pathExists(fpath) {
		err :=os.MkdirAll(fpath,0755)
		if err != nil {
			golog.Error("文件夹创建失败","path",fpath,"err",err)
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	resp, err := c.keysAPI.Get(ctx, c.prefix, clientv3.WithPrefix())
	cancel()
	if err != nil {
		golog.Error("获取etcd key 失败", "key",c.prefix,"err", err)
		return
	}

	var backupmap = make(map[string]interface{})

	for _, kv := range resp.Kvs {
		path := strings.Split(string(kv.Key), "/")
		lastKey := path[len(path)-1]

		deepestMap := deepSearch(backupmap, path[0:len(path)-1])

		deepestMap[lastKey] = string(kv.Value)
	}

	filename := filepath.Join(fpath,c.prefix)+ "_" + time.Now().Format(shortDate)+".json"
	getOperating().SetConfigFile(filename)
	err = getOperating().WriteConfig(backupmap)
	if err != nil {
		golog.Error("备份写入文件失败", "file",filename,"err", err)
		return
	}

	golog.Info("etcd 备份成功","file",filename)

}

func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			// intermediate key does not exist
			// => create it and continue from there
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]interface{})
		if !ok {
			// intermediate key is a value
			// => replace with a new map
			m3 = make(map[string]interface{})
			m[k] = m3
		}
		// continue search from here
		m = m3
	}
	return m
}


func pathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
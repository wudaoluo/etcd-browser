package model

import (
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/kubernetes/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
	"github.com/syndtr/goleveldb/leveldb/util"
	e "github.com/wudaoluo/etcd-browser"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"sort"
	"time"
)

const shortForm = "2006-01-02 15:04:05"

type Record struct {
	Key     string
	Value   string
	Version int
	Time    string
	Type    string
}

type Records []*Record

func (p Records) Len() int           { return len(p) }
func (p Records) Less(i, j int) bool { return p[i].Version > p[j].Version }
func (p Records) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var cnf = e.GetConfigInstance()

//etcd key ? isDir :{key == false, dir == true}
func Get(key string, isDir bool) []*Record {
	var rs []*Record
	var err error

	iter := db.NewIterator(util.BytesPrefix([]byte(key)), nil)

	var i = 0
	for iter.Next() {
		//返回的最大数量
		if i >= cnf.GetInt("history_num") {
			break
		}

		var r = &Record{}
		err = json.Unmarshal(iter.Value(), r)
		if err != nil {
			golog.Error("json解码失败,丢掉这条数据", "err", err)
			continue
		}

		if isDir && r.Type != etcdlib.WATCH_EVENT_DELETE {
			continue
		}


		rs = append(rs, r)
		i++
	}

	sort.Sort(Records(rs))
	return rs
}

func Put(key, reallyKey, value []byte, version int64, t string) error {
	v, err := json.Marshal(&Record{
		Key:     string(reallyKey),
		Value:   string(value),
		Version: int(version),
		Time:    time.Now().Format(shortForm),
		Type:    t,
	})

	if err != nil {
		return err
	}

	if cnf.GetBool("history_really_del") && t == etcdlib.WATCH_EVENT_DELETE {
		//将leveldb中所有 reallyKey 信息删除，但上一级会保留删除时该key的配置
		Del(reallyKey)
	}

	return db.Put(key, v, nil)
}

func Del(key []byte) {
	iter := db.NewIterator(util.BytesPrefix([]byte(key)), nil)
	for iter.Next() {
		err := db.Delete(iter.Key(), nil)
		golog.Warn("删除key", "key", string(iter.Key()))
		if err != nil {
			golog.Error("删除数据失败", "key", string(iter.Key()), "err", err)
		}
	}
}

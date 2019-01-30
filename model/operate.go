package model

import (
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/kubernetes/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
	"github.com/syndtr/goleveldb/leveldb/util"
	"sort"
	"time"
)

const shortForm = "2006-01-01 15:04:05"

type Record struct {
	Value string  `json:"value"`
	Version int   `json:"version"`
	Time string   `json:"time"`
}

type Records []*Record

func (p Records) Len() int           { return len(p) }
func (p Records) Less(i, j int) bool { return p[i].Version > p[j].Version }
func (p Records) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func Get(key string) []*Record{
	var rs []*Record
	var err error

	iter := db.NewIterator(util.BytesPrefix([]byte(key)),nil)

	for iter.Next() {
		var r = &Record{}
		err = json.Unmarshal(iter.Value(),r)
		if err != nil {
			golog.Error("json解码失败,丢掉这条数据","err",err)
			continue
		}
		rs = append(rs,r)
	}

	sort.Sort(Records(rs))
	return rs
}


func Put(key,value []byte,version int64) error {
	v ,err:= json.Marshal(&Record{
		Value:string(value),
		Version:int(version),
		Time:time.Now().Format(shortForm),
	})
	if err != nil {
		return err
	}

	return db.Put(key,v,nil)
}


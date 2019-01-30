package model

import (
	"context"
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB



func DBInit(ctx context.Context) {
	//cnf:= e.GetConfigInstance()
	//cnf.GetString("db.source_name")
	var err error
	db, err = leveldb.OpenFile(".etcd_browser" , nil)
	if err != nil {
		golog.Fatalw("leveldb 打开失败","err",err)
		panic(err)
	}
	var s = &leveldb.DBStats{}
	err = db.Stats(s)
	if err != nil {
		golog.Fatalw("'leveldb stats 获取失败","err",err)
		panic(err)
	}

	go func() {
		select {
		case <-ctx.Done():
			golog.Debug("关闭leveldb ...")
			DBclose()
		}
	}()
}


func DBclose() {
	err := db.Close()
	if err != nil {
		golog.Errorw("leveldb 关闭失败", "err", err)
	}

}




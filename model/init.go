package model

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
	//"github.com/ThreeKing2018/goutil/golog"
	e "github.com/wudaoluo/etcd-browser"

)

var db *sql.DB


//username:password@tcp(dbhost:dbport)/dbname?charset=utf8
func Init() {
	//var err error

	cnf:= e.GetConfigInstance()

	fmt.Println("aaaaaaaaaaaaaaa",cnf.GetString("db.source_name"))
	//db, err := sql.Open("sqlite3", cnf.GetString("db.source_name"))
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//golog.Info("mysql连接成功")
	//
	//_,err=db.Exec("SHOW BC")
	//fmt.Println("err",err)
	//initService()
}


//设置数据库前缀
func tableName(name string) string {
	return name
}


//
//
//var DBuser *userService
//
//
//func initService() {
//	DBuser = &userService{}
//}

/*
先判断文件是否存在 不存在进行初始化
  //创建表//delete from BC;，SQLite字段类型比较少，bool型可以用INTEGER，字符串用TEXT
    sqlStmt := `create table BC (b_code text not null primary key, c_code text not null, code_type INTEGER, is_new INTEGER);`
    _, err = db.Exec(sqlStmt)
    if err != nil {
        fmt.Println("create table error->%q: %s\n", err, sqlStmt)
        return
    }
    //创建索引，有索引和没索引性能差别巨大，根本就不是一个量级，有兴趣的可以去掉试试
    _, err = db.Exec("CREATE INDEX inx_c_code ON BC(c_code);")
    if err != nil {
        fmt.Println("create index error->%q: %s\n", err, sqlStmt)
        return
    }
*/
package v3

import (
	e "github.com/wudaoluo/etcd-browser"
	"github.com/wudaoluo/etcd-browser/etcdlib"
)


func init() {
	/*
	    "etcd_root_key": "root",
    "etcd_addr": "127.0.0.1:2379",
	*/
	cnf:= e.GetConfigInstance()
	etcdlib.SetEtcd(cnf.GetStringSlice("etcd_addr"),
		cnf.GetString("etcd_root_key"))
}
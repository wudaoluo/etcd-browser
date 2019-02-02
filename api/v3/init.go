package v3

import (
	"context"
	"crypto/tls"
	"github.com/ThreeKing2018/goutil/golog"
	e "github.com/wudaoluo/etcd-browser"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"github.com/wudaoluo/etcd-browser/model"
	"go.etcd.io/etcd/pkg/transport"
	"gopkg.in/robfig/cron.v2"
)

const ETCD_V2 = "v2"


func Init(ctx context.Context) {
	cnf:= e.GetConfigInstance()
	if  cnf.GetString("etcd_version") == ETCD_V2 {
		return
	}
	var tls *tls.Config

	if cnf.GetString("cert_file") != "" &&
		cnf.GetString("key_file") != "" &&
		cnf.GetString("ca_file") != "" {

		var err error
		tlsInfo := transport.TLSInfo{
			CertFile:      cnf.GetString("cert_file"),
			KeyFile:       cnf.GetString("key_file"),
			TrustedCAFile: cnf.GetString("ca_file"),
		}
		tls, err = tlsInfo.ClientConfig()
		if err != nil {
			panic(err)
		}

	}

	etcdlib.SetEtcd(ctx,
		cnf.GetStringSlice("etcd_addr"),
		cnf.GetString("etcd_root_key"),
		tls)


	//etcd 备份
	c := cron.New()
	_, err := c.AddFunc(cnf.GetString("etcd_back_time"), func() {
		//当etcd_back_path没有设置则在程序运行目录创建备份文件
		etcdlib.Backup(cnf.GetString("etcd_back_path"))
	})
	if err != nil {
		//每次启动时候都要去定一下日志 查看有没有启动成功
		golog.Fatal("启动定时任务失败")
		return
	}
	c.Start()

	//后悔药
	if cnf.GetBool("etcd_watch") {
		etcdlib.Watch(model.Put)
	}
}
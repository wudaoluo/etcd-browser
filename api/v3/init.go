package v3

import (
	"context"
	"crypto/tls"
	e "github.com/wudaoluo/etcd-browser"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"github.com/wudaoluo/etcd-browser/model"
	"go.etcd.io/etcd/pkg/transport"
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
	etcdlib.SetEtcd(ctx,cnf.GetStringSlice("etcd_addr"),
		cnf.GetString("etcd_root_key"),tls)

	if cnf.GetBool("etcd_watch") {
		etcdlib.Watch(model.Put)
	}
}
package v3

import (
	"crypto/tls"
	e "github.com/wudaoluo/etcd-browser"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"go.etcd.io/etcd/pkg/transport"
)


func init() {
	cnf:= e.GetConfigInstance()

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
	etcdlib.SetEtcd(cnf.GetStringSlice("etcd_addr"),
		cnf.GetString("etcd_root_key"),tls)
}
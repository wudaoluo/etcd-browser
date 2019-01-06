package etcdlib

import (
	"context"
	"crypto/tls"
)

var EtcdClient Clienter

func SetEtcd(endpoint []string, Prefix string,tls *tls.Config) {
	var err error
	EtcdClient, err = New(endpoint, Prefix,tls)
	if err != nil {
		panic(err)
	}
}

func Get(key string) (*Node, error) {
	return EtcdClient.Get(key)
}

func GetContext(ctx context.Context, key string) (*Node, error) {
	return EtcdClient.GetContext(ctx, key)
}

func List(key string) ([]*Node, error) {
	return EtcdClient.List(key)
}

func ListContext(ctx context.Context, key string) ([]*Node, error) {
	return EtcdClient.ListContext(ctx, key)
}

func Put(key, value string) error {
	return EtcdClient.Put(key, value)
}

func PutContext(ctx context.Context, key, value string) error {
	return EtcdClient.PutContext(ctx, key, value)
}

func Create(key, value string) error {
	return EtcdClient.Create(key, value)
}

func CreateContext(ctx context.Context, key, value string) error {
	return EtcdClient.CreateContext(ctx, key, value)
}

func CreateDir(key string) error {
	return EtcdClient.CreateDir(key)
}

func CreateDirContext(ctx context.Context, key string) error {
	return EtcdClient.CreateDirContext(ctx, key)
}

func Delete(key string) error {
	return EtcdClient.Delete(key)
}

func DeleteContext(ctx context.Context, key string) error {
	return EtcdClient.DeleteContext(ctx, key)
}

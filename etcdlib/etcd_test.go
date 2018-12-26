package etcdlib

import (
	"fmt"
	"testing"
)

const (
	TEST_ETCD_ADDR = "127.0.0.1:2379"
	TEST_ROOT_KEY  = "root"
)


func init() {
	SetEtcd([]string{TEST_ETCD_ADDR},TEST_ROOT_KEY)
}

//func TestCreate(t *testing.T) {
//	err :=EtcdClient.Create("/a","a1")
//	if err != nil {
//		t.Error(err)
//	}
////}
//
//func TestPut(t *testing.T) {
//	err := EtcdClient.Put("/a","b1",false)
//		if err != nil {
//			t.Error(err)
//		}
//}
//
//func TestGet(t *testing.T) {
//	n,err :=EtcdClient.Get("/a")
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(n)
//}
//
//func TestDelete(t *testing.T) {
//	err :=EtcdClient.Delete("/a")
//	if err != nil {
//		t.Error(err)
//	}
//}

func TestList(t *testing.T) {
	n,err :=EtcdClient.List("/")
	if err != nil {
		t.Error(err)
	}

	for _,i := range n {
		fmt.Println(i)
	}
}
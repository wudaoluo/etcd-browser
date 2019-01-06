package etcdlib

import (
	"fmt"
	"testing"
)


func init() {
	SetEtcd([]string{"127.0.0.1:32777","127.0.0.1:32778","127.0.0.1:32779"}, "root/ddsds")
}


func Testget(t *testing.T) {
	r ,err :=EtcdClient.Get("jjj123")
	fmt.Println(r)
	fmt.Println(err)
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
//		}l
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
//
//func TestPut(t *testing.T) {
//	//_,err := Get("/test/sddss1/aaa")
//
//	n, err := Get("/tesaaaaa/dsdssdsdds")
//	fmt.Println(err, "err")
//	fmt.Println("n", n)
//	return
//
//	keys := strings.Split("/test/sddss1/aaa/sss/dssd/sasa/sassaas", "/")
//	fmt.Println(keys)
//
//	root := "/"
//	for _, key := range keys {
//		if key == "" {
//			continue
//		}
//
//		root = path.Join(root, key)
//
//		if key == keys[len(keys)-1] {
//			Create(root, "a")
//			return
//		}
//		_, err := Get(root)
//
//		if err != nil {
//			fmt.Println(root)
//			err = EtcdClient.CreateDir(root)
//			if err != nil {
//				t.Error(err)
//			}
//		}
//
//	}

	//err := EtcdClient.Put("/test/sddss1/aaa/sss/dssd","aa")
	//if err != nil {
	//	t.Error(err)
	//}

//}
//
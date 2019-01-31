package main

import (
	"fmt"
	"github.com/wudaoluo/etcd-browser/model"
	"context"

)
func main() {
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	model.DBInit(ctx)
	//err := model.Put([]byte("/test1/sddss1/aaa/sss/dssd/sasa/sassaas"),[]byte("a"),1223)
	//fmt.Println(err)

	r := model.Get("root/test1/sddss1/aaa/sss/dssd/sasa/sassaas")
	fmt.Println(r)

	for _,i := range r {
		fmt.Println(i.Key)
	}
}

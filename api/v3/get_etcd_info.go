package v3

import (
	"github.com/emicklei/go-restful"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
)

func Leader(request *restful.Request, response *restful.Response) {
	member, err :=etcdlib.MembersHandler()
	if err != nil {
		_ = response.WriteError(http.StatusNotFound,err)
		return
	}
	_ = response.WriteEntity(member)
}

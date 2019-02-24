package v3

import (
	"github.com/emicklei/go-restful"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
	"path"
)

func PutKeys(request *restful.Request, response *restful.Response) {
	key :=  path.Join("/",request.PathParameter("subpath"))
	value := request.QueryParameter("value")
	err := etcdlib.Put(key, value)
	if err != nil {
		response.WriteError(http.StatusInternalServerError,err)
		return
	}

	response.WriteEntity(nil)
}

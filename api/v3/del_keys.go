package v3

import (
	"github.com/emicklei/go-restful"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
	"path"
)

func DelKeys(request *restful.Request, response *restful.Response)  {
	key :=  path.Join("/",request.PathParameter("subpath"))
	err := etcdlib.Delete(key)
	if err != nil {
		response.WriteError(http.StatusInternalServerError,err)
		return
	}

	response.WriteEntity(nil)
}

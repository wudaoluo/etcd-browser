package v3

import (
	"github.com/emicklei/go-restful"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"github.com/wudaoluo/etcd-browser/model"
	"net/http"
	"path"
)


type respHistoryValue struct {
	Action string `json:"action"`
	Node []*model.Record `json:"node"`
	Key string `json:"key"`

}

func History(request *restful.Request, response *restful.Response) {
	key := path.Join("/",request.PathParameter("subpath"))
	etcdKey, _, err := etcdlib.EnsureKey(key)
	if err != nil {
		response.WriteError(http.StatusInternalServerError,err)
	}

	record := model.Get(etcdKey)
	response.WriteEntity(respHistoryValue{Action:"history",Node:record,Key:key})
}


func Restore(request *restful.Request, response *restful.Response) {
	key := path.Join("/",request.PathParameter("subpath"))
	value := request.QueryParameter("value")
	err := etcdlib.Put(key, value)
	if err != nil {
		response.WriteError(http.StatusInternalServerError,err)
		return
	}

	response.WriteEntity(nil)
}
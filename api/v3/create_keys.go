package v3

import (
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/emicklei/go-restful"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
	"path"
	"strings"
)

type respPostKeysValue struct {
	Action string
}

func PostKeys(request *restful.Request, response *restful.Response) {
	key := path.Join("/",request.PathParameter("subpath"))
	isDir := stringToBool(request.QueryParameter("dir"))

	keys := strings.Split(key, "/")
	rootKey := "/"
	keysLen := len(keys)
	var err error

	for i := 0; i < keysLen; i++ {
		if keys[i] == "" {
			continue
		}

		rootKey = path.Join(rootKey, keys[i])
		if i == (keysLen-1) && ! isDir {
			value := request.QueryParameter("value")
			err = etcdlib.Create(rootKey, value)
			if err != nil {
				golog.Errorf("etcdlib.Create(rootKey, value)","key",rootKey,"value",value,"err",err)
				response.WriteError(http.StatusInternalServerError,err)
				return
			}
			break
		}

		_, err = etcdlib.Get(rootKey)
		if err == etcdlib.ErrorKeyNotFound {
			err = etcdlib.CreateDir(rootKey)
			if err == etcdlib.ErrorPutKey {
				continue
			}
			if err != nil {
				golog.Errorf("etcdlib.CreateDir(rootKey)","rootKey",rootKey,"err",err)
				response.WriteError(http.StatusInternalServerError,err)
				return
			}
			continue
		}

		if err != nil {
			golog.Errorf("etcdlib.Get(rootKey)","rootKey",rootKey,"err",err)
			response.WriteError(http.StatusInternalServerError,err)
			return
		}

	}

	response.WriteEntity(respPostKeysValue{Action:"set"})
}


func stringToBool(a string) bool {
	if strings.ToLower(a)=="true" {
		return true
	}

	return false
}



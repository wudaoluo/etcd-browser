package v3

import (
	"github.com/emicklei/go-restful"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
	"path"
)

type Node struct {
	Key   string  `json:"key"`
	Value string  `json:"value,omitempty"`
	IsDir bool    `json:"dir,omitempty"`
	Nodes []*Node `json:"nodes,omitempty"`
}

func parseNode(node *etcdlib.Node) *Node {
	return &Node{
		Key:   string(node.Key),
		Value: string(node.Value),
		IsDir: node.IsDir,
	}
}

type respKeysValue struct {
	Action string `json:"action"`
	Node *Node `json:"node"`
}

func Keys(request *restful.Request, response *restful.Response)  {

	key := path.Join("/",request.PathParameter("subpath"))

	node, err := etcdlib.Get(key)
	if err != nil {
		response.WriteError(http.StatusNotFound,err)
		return
	}
	if node.IsDir {

		nodes, err := etcdlib.List(key)
		if err != nil {
			response.WriteError(http.StatusInternalServerError,err)
			return
		}

		realNodes := &Node{
			Key:   key,
			IsDir: true,
		}
		for _, node := range nodes {
			realNodes.Nodes = append(realNodes.Nodes, parseNode(node))
		}
		response.WriteEntity(respKeysValue{Action:"get",Node:realNodes})
	} else {
		/*
			node, err := client.Get(key)
			if err != nil {
				return nil, err
			}

			return parseNode(node), nil
		*/
		response.WriteEntity(respKeysValue{Action:"get",Node:parseNode(node)})
	}
}

package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
)


/*
"key":"/aa/aaas",
"dir":true,
"modifiedIndex":21,
"value":"asss",
"createdIndex":21
*/
type Node struct {
	Key           string `json:"key"`
	Value         string `json:"value,omitempty"`
	IsDir         bool   `json:"dir,omitempty"`
	//ModifiedIndex int64  `json:"modifiedIndex"`
	//CreatedIndex  int64  `json:"createdIndex"`
	Nodes         []*Node `json:"nodes,omitempty"`
}


func parseNode(node *etcdlib.Node) *Node {
	return &Node{
		Key:   string(node.Key),
		Value: string(node.Value),
		IsDir: node.IsDir,
	}
}
const (
	TEST_ETCD_ADDR = "127.0.0.1:2379"
	TEST_ROOT_KEY  = "root"
)


func init() {
	etcdlib.SetEtcd([]string{TEST_ETCD_ADDR}, TEST_ROOT_KEY)
}

func Keys(c *gin.Context) {

	key := c.Param("action")
	nodes,err :=etcdlib.List(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":err})
	}

	realNodes := &Node{
		Key:key,
		IsDir:true,
	}
	for _, node := range nodes {
		realNodes.Nodes = append(realNodes.Nodes, parseNode(node))
	}
	c.JSON(http.StatusOK,gin.H{"action":"get","node":realNodes})
}


func Root(c *gin.Context) {
	key := "/"
	nodes,err :=etcdlib.List(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"err":err})
	}

	realNodes := &Node{
		Key:key,
		IsDir:true,
	}
	for _, node := range nodes {
		realNodes.Nodes = append(realNodes.Nodes, parseNode(node))
	}
	c.JSON(http.StatusOK,gin.H{"action":"get","node":realNodes})
}
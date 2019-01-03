package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
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



func Keys(c *gin.Context) {

	key := c.Param("action")

	node, err := etcdlib.Get(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	if node.IsDir {

		nodes, err := etcdlib.List(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			return
		}

		realNodes := &Node{
			Key:   key,
			IsDir: true,
		}
		for _, node := range nodes {
			realNodes.Nodes = append(realNodes.Nodes, parseNode(node))
		}
		c.JSON(http.StatusOK, gin.H{"action": "get", "node": realNodes})
	} else {
		/*
			node, err := client.Get(key)
			if err != nil {
				return nil, err
			}

			return parseNode(node), nil
		*/
		c.JSON(http.StatusOK, gin.H{"action": "get", "node": parseNode(node)})
	}
}

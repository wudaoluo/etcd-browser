package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
)

func PutKeys(c *gin.Context) {
	key := c.Param("action")
	_ ,isDir := c.GetQuery("dir")
	realNodes := &Node{}
	if isDir {
		etcdlib.CreateDir(key)
		realNodes.Key = key
		realNodes.IsDir = isDir


	}else {
		value := c.Query("value")
		etcdlib.Create(key,value)
		realNodes.Key = key
		realNodes.Value = value
	}
	c.JSON(http.StatusOK,gin.H{"action":"set","node":realNodes})
}

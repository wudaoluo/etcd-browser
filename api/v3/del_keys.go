package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
)

func DelKeys(c *gin.Context) {
	key := c.Param("action")
	_, isDir := c.GetQuery("dir")

	etcdlib.Delete(key)

	prevNodes := &Node{Key:key}
	realNodes := &Node{
		Key:key,
	}
	if isDir {
		realNodes.IsDir = isDir
		prevNodes.IsDir = isDir
		//删除目录
	}
	c.JSON(http.StatusOK,gin.H{"action":"delete","node":realNodes,"prevNode":prevNodes})
}

package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
)

func Leader(c *gin.Context) {
	member, err :=etcdlib.MembersHandler()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK,member)
}

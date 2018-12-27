package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
)

func PutKeys(c *gin.Context) {
	key := c.Param("action")
	value := c.Query("value")
	err := etcdlib.Put(key, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

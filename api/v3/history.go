package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"github.com/wudaoluo/etcd-browser/model"
	"net/http"
)


func History(c *gin.Context) {
	key := c.Param("action")
	etcdKey, _, err := etcdlib.EnsureKey(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	record := model.Get(etcdKey)
	c.JSON(http.StatusOK, gin.H{"action": "history", "node": record,"key":key})
}


func Restore(c *gin.Context) {
	key := c.Param("action")
	value := c.Query("value")
	err := etcdlib.Create(key, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
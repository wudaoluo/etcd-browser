package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/wudaoluo/etcd-browser/etcdlib"
	"net/http"
	"path"
	"strings"
)

func PostKeys(c *gin.Context) {
	key := c.Param("action")
	_, isDir := c.GetQuery("dir")
	keys := strings.Split(key, "/")
	rootKey := "/"
	keysLen := len(keys)
	var err error

	for i := 0; i < keysLen; i++ {
		if keys[i] == "" {
			continue
		}

		rootKey = path.Join(rootKey, keys[i])
		if i == (keysLen-1) && !isDir {
			value := c.Query("value")
			err = etcdlib.Create(rootKey, value)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error(), "key": rootKey})
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
				c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error(), "key": rootKey})
				return
			}
			continue
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error(), "key": rootKey})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"action": "set"})
}



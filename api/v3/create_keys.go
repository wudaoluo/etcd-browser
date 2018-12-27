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
		if keys[i] == keys[keysLen-1] && !isDir {
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

/*

func TestPut(t *testing.T) {
	_,err := Get("/test/sddss1/aaa")
	fmt.Println(err,"err")

	keys := strings.Split("/test/sddss1/aaa/sss/dssd/sasa/sassaas","/")
	fmt.Println(keys)

	root := "/"
	for _,key := range keys {
		if key == "" {
			continue
		}


		root = path.Join(root,key)

		if key == keys[len(keys)-1] {
			Create(root,"a")
			return
		}
		_ ,err:=Get(root)


		if err != nil  {
			fmt.Println(root)
			err =EtcdClient.CreateDir(root)
			if err != nil {
				t.Error(err)
			}
		}

	}

	//err := EtcdClient.Put("/test/sddss1/aaa/sss/dssd","aa")
	//if err != nil {
	//	t.Error(err)
	//}


}

*/

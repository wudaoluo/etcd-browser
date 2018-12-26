package main

import (
	"github.com/gin-gonic/gin"
	apiv3 "github.com/wudaoluo/etcd-browser/api/v3"
)


func main() {


	router := gin.Default()

	v3 :=router.Group("v2")
	v3.GET("/keys/*action", apiv3.Keys)  //查询
	v3.PUT("/keys/*action", apiv3.PutKeys)
	v3.DELETE("/keys/*action", apiv3.DelKeys)
	//v3.GET("/keys/", apiv3.Root)
	//v3.GET("/keys/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":3000")

}

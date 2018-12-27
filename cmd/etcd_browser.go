package main

import (
	"github.com/gin-gonic/gin"
	apiv2 "github.com/wudaoluo/etcd-browser/api/v2"
	apiv3 "github.com/wudaoluo/etcd-browser/api/v3"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	router.StaticFile("/index", "../frontend/index.html")
	router.StaticFile("/main.css", "../frontend/main.css")
	router.StaticFile("/etcdbrowser.js", "../frontend/etcdbrowser.js")
	router.StaticFile("/favicon.ico", "../frontend/favicon.ico")
	router.Static("/angular-xeditable", "../frontend/angular-xeditable")

	//etcd version 2
	v2 := router.Group("/v2")
	v2.Any("/*action", apiv2.ReverseProxy)

	//etcd version 3
	v3 := router.Group("/v3")
	v3.GET("/keys/*action", apiv3.Keys)
	v3.POST("/keys/*action", apiv3.PostKeys)
	v3.DELETE("/keys/*action", apiv3.DelKeys)
	v3.PUT("/keys/*action", apiv3.PutKeys)

	router.Run(":3000")

}

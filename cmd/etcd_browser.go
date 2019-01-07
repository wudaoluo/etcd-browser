package main

import (
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/gin-gonic/gin"
	e "github.com/wudaoluo/etcd-browser"
	apiv2 "github.com/wudaoluo/etcd-browser/api/v2"
	apiv3 "github.com/wudaoluo/etcd-browser/api/v3"
	"net/http"
	"time"
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

	//v3.GET("/stats/store",apiv3.Leader)
	//v3.GET("/stats/leader",apiv3.Leader)
	v3.GET("/stats/self",apiv3.Leader)

	cnf:= e.GetConfigInstance()

	httpServer := &http.Server{
		Addr:           cnf.GetString("listen"),
		Handler:        router,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		IdleTimeout:    30 * time.Second,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		golog.Fatal("start http fail:", err.Error())
	}


}

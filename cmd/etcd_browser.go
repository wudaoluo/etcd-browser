package main

import (
	"context"
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/gin-gonic/gin"
	e "github.com/wudaoluo/etcd-browser"
	apiv2 "github.com/wudaoluo/etcd-browser/api/v2"
	apiv3 "github.com/wudaoluo/etcd-browser/api/v3"
	"github.com/wudaoluo/etcd-browser/model"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx,cancel := context.WithCancel(context.Background())

	model.DBInit(ctx)
	apiv3.Init(ctx)


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
	v3.GET("/stats/self",apiv3.Leader)
	v3.POST("/history/*action",apiv3.History)
	v3.PUT("/restore/*action",apiv3.Restore)


	cnf:= e.GetConfigInstance()

	httpServer := &http.Server{
		Addr:           cnf.GetString("listen"),
		Handler:        router,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		IdleTimeout:    30 * time.Second,
	}


	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt,os.Kill, syscall.SIGTERM)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := httpServer.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			golog.Error("HTTP server Shutdown: %v", err)
		}
		cancel()
	}()


	err := httpServer.ListenAndServe()
	if err != nil {
		golog.Error("http err:", err.Error())
	}
}

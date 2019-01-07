package v2

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	e "github.com/wudaoluo/etcd-browser"
)

func ReverseProxy(ctx *gin.Context) {
	cnf:= e.GetConfigInstance()
	var targets = []*url.URL{}
	for _, etcd_addr := range cnf.GetStringSlice("etcd_addr") {
		targets = append(targets,&url.URL{
			Scheme: "http",
			Host:   etcd_addr,
		})
	}


	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}

	//proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

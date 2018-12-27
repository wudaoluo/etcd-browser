package v2

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(ctx *gin.Context) {
	var targets = []*url.URL{
		{
			Scheme: "http",
			Host:   "127.0.0.1:12379",
		},
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

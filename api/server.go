package api

import (
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/emicklei/go-restful"
	"github.com/spf13/cobra"
	"net/http"
	"path"
	apiv3 "github.com/wudaoluo/etcd-browser/api/v3"
)


type argStruct struct {
	version bool
	configfile string
	//debug bool
	//logdir string
}

var Arg = new(argStruct)


func NewServerCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:"etcd_browser",
		Short:"bbbbbbbbbbb",
		Long:"aaaaaaaaaaaaaaaa",
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
		
	}
/*
	flag.BoolVar(&Arg.version,"version",false,"打印版本号")
	//flag.BoolVar(&Arg.debug,"debug",true,"open debug default false")
	flag.StringVar(&Arg.configfile,"c",DEFAULT_CONFIG_NAME,"specify config file")
*/
	rootCmd.Flags().StringVar(
		&Arg.configfile,"c","config.toml","config file (default is config.toml)")

	rootCmd.Flags().BoolVar(
		&Arg.version,"version",false,"print version")
	return rootCmd
}


//Container > webservice > route
//https://blog.csdn.net/Daniel_greenspan/article/details/78624725
func Run() {
	static := new(restful.WebService)
	restful.Add(static)

	static.Route(static.GET("/").To(index))
	static.Route(static.GET("/index.html").To(staticFile))
	static.Route(static.GET("/main.css").To(staticFile))
	static.Route(static.GET("/etcdbrowser.js").To(staticFile))
	static.Route(static.GET("/favicon.ico").To(staticFile))
	static.Route(static.GET("/angular-xeditable/{subpath:*}").To(staticFromPathParam))

	v3 := new(restful.WebService)
	v3.Path("/v3")

	// Consumers 响应的 Content-Type
	//v3.Consumes(restful.MIME_JSON)

	// Produces 请求的 Accept
	//v3.Produces(restful.MIME_JSON)
	restful.Add(v3)


	v3.Route(v3.GET("/keys/").To(apiv3.Keys))
	v3.Route(v3.GET("/keys/{subpath:*}").To(apiv3.Keys))
	v3.Route(v3.POST("/keys/{subpath:*}").To(apiv3.PostKeys))
	v3.Route(v3.DELETE("/keys/{subpath:*}").To(apiv3.DelKeys))
	v3.Route(v3.PUT("/keys/{subpath:*}").To(apiv3.PutKeys))
	v3.Route(v3.GET("/stats/self").To(apiv3.Leader))
	v3.Route(v3.POST("/history/{subpath:*}").To(apiv3.History))
	v3.Route(v3.PUT("restore/{subpath:*}").To(apiv3.Restore))




	golog.Info("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: restful.DefaultContainer}
	golog.Fatal(server.ListenAndServe())

}


const rootdir = "../frontend"


func index(req *restful.Request, resp *restful.Response) {
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		path.Join(rootdir, "index.html"))
}


func staticFile(req *restful.Request, resp *restful.Response) {
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		path.Join(rootdir, req.SelectedRoutePath()))
}

func staticFromPathParam(req *restful.Request, resp *restful.Response) {
	actual := path.Join(rootdir, "angular-xeditable",req.PathParameter("subpath"))
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actual)
}

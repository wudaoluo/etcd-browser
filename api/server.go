package api

import (
	"fmt"
	"github.com/ThreeKing2018/goutil/golog"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
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

	rootCmd.AddCommand(getVersion())
	return rootCmd
}


func getVersion() *cobra.Command {
	version := &cobra.Command{
		Use: "version",
		Short:"打印版本号",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("etcd-browser v1.0")
		},
	}

	return version

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
	static.Route(static.GET("/apidocs/").To(staticSwagger))
	static.Route(static.GET("/apidocs/{subpath:*}").To(staticSwagger))


	v3 := new(restful.WebService)
	v3.Path("/v3")

	// Consumers 响应的 Content-Type
	//v3.Consumes(restful.MIME_JSON)

	// Produces 请求的 Accept
	//v3.Produces(restful.MIME_JSON)
	restful.Add(v3)


	v3.Route(v3.GET("/keys/").
		To(apiv3.Keys).
		Doc("获取key的value"))
	v3.Route(v3.GET("/keys/{subpath:*}").To(apiv3.Keys).
		Doc("aaaa"))
	v3.Route(v3.POST("/keys/{subpath:*}").To(apiv3.PostKeys))
	v3.Route(v3.DELETE("/keys/{subpath:*}").To(apiv3.DelKeys))
	v3.Route(v3.PUT("/keys/{subpath:*}").To(apiv3.PutKeys))
	v3.Route(v3.GET("/stats/self").To(apiv3.Leader))
	v3.Route(v3.POST("/history/{subpath:*}").To(apiv3.History))
	v3.Route(v3.PUT("restore/{subpath:*}").To(apiv3.Restore))


	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))


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

func staticSwagger(req *restful.Request, resp *restful.Response) {
	actual := path.Join(rootdir, "dist",req.PathParameter("subpath"))
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actual)
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Etcd-Browser",
			Description: "etcd v3 webIU api接口",
			Contact: &spec.ContactInfo{
				Name:  "carlo",
				URL:   "https://github.com/wudaoluo/etcd-browser",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "0.2.1",
		},
	}
	//swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
	//	Name:        "users",
	//	Description: "Managing users"}}}
}

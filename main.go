package main

import (
	"embed"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"jamma/api"
	"jamma/chat"
	"jamma/config"
	_ "jamma/docs"
	"net/http"
)

//go:embed all:www
var wwwFiles embed.FS

// @title 接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/
// @InstanceName jamma
// @query.collection.format multi

func main() {

	config.Load()

	//原本的Main函数
	engine := web.CreateEngine()
	engine.Static("/static", "static")
	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//聊天室
	chat.Register(engine.Group("/chat"))

	//注册接口文档
	web.RegisterSwaggerDocs(&engine.RouterGroup, "jamma")

	//附件
	engine.Static("/attach", "attach")

	//注册静态页面
	fs := engine.FileSystem()
	fs.Put("", http.FS(wwwFiles), "www", "index.html")

	//启动
	engine.Serve()
}

package main

import (
	"embed"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"jamma/api"
	"jamma/chat"
	"jamma/config"
	_ "jamma/docs"
	"jamma/types"
	"net/http"
)

//go:embed all:www
var wwwFiles embed.FS

// @title jamma接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/
// @InstanceName jamma
// @query.collection.format multi

func main() {

	config.Load()

	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	db.Engine.Sync2(new(types.Email), new(types.Game),
		new(types.Box), new(types.Recharge),
		new(types.SignIn), new(types.User),
		new(types.Me), new(types.Password), new(types.UserHistory),
		new(types.ChatHongBao), new(types.ChatQiangHongBao))
	//原本的Main函数
	engine := web.CreateEngine()
	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//聊天室
	chat.Register(engine.Group("/chat"))

	//注册接口文档
	web.RegisterSwaggerDocs(&engine.RouterGroup, "jamma")

	//附件
	//engine.Static("/attach", "attach")
	//静态文件
	engine.Static("/static", "static")
	//注册静态页面
	fs := engine.FileSystem()
	fs.Put("", http.FS(wwwFiles), "www", "index.html")

	//启动
	engine.Serve()
}

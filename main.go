package main

import (
	"arcade/api"
	"arcade/chat"
	"arcade/config"
	"arcade/types"
	"arcade/weixin"
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

//go:embed all:www
var wwwFiles embed.FS

// @title arcade接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/
// @InstanceName arcade
// @query.collection.format multi

func main() {
	config.Load()

	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Engine.Sync2(
		new(types.User), new(types.Me), new(types.Password), new(types.UserHistory),
		new(types.Game), new(types.Box), new(types.Recharge), new(types.Exchange),
		new(types.SignIn), new(types.Email),
		new(chat.RedPacket), new(chat.GrabPacket))
	if err != nil {
		log.Fatal(err)
	}

	//微信接口
	weixin.Open()

	//原本的Main函数
	engine := web.CreateEngine()
	engine.Use(api.Cors())
	engine.Use(api.CatchError)

	//注册前端接口
	api.RegisterRoutes(engine.Group("/api"))

	//聊天室
	chat.Register(engine.Group("/chat"))

	//注册接口文档
	web.RegisterSwaggerDocs(&engine.RouterGroup, "arcade")

	//附件
	//engine.Static("/attach", "attach")
	//静态文件
	//engine.Static("/static", "static")

	//注册静态页面
	fs := engine.FileSystem()
	//fs.Put("", http.FS(wwwFiles), "www", "index.html")
	fs.Put("", gin.Dir("www", false), "", "index.html")

	//engine.Static("/", "www")

	//启动
	go engine.Serve()

	/////////////////////////
	/////改用LetsEncrypt//////
	////////////////////////

	//初始化autocert
	manager := &autocert.Manager{
		Cache:      autocert.DirCache("certs"),
		Email:      "jason@zgwit.com",
		HostPolicy: autocert.HostWhitelist("gamebox.zgwit.cn"),
		Prompt:     autocert.AcceptTOS,
	}

	//创建server
	svr := &http.Server{
		Addr:      "0.0.0.0:443",
		TLSConfig: manager.TLSConfig(),
		Handler:   engine,
	}

	//监听https
	err = svr.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}

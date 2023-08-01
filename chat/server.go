package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func Register(router *gin.RouterGroup) {

	server := socketio.NewServer(nil)

	// 建立连接
	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("连接成功：", s.ID())

		fmt.Println("query参数：", s.URL().RawQuery)
		s.Emit("message", gin.H{"status": 200, "data": "123"})
		s.Join("chat1")
		server.BroadcastToRoom("/", "caht1", "notice", "通知")
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("连接错误:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		s.LeaveAll()
		if uid := s.Context(); uid != nil {
			log.Printf("用户[%s]断开连接", uid)
		}
		log.Println("关闭连接：", s.ID(), reason)
	})

	// 服务/事件
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice收到内容：:", msg)
		s.Emit("notice", "have "+msg) // 向client回复内容
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn, msg string) string {
		last := s.Context().(string)
		log.Println(last)
		log.Println("msg", msg)
		s.Emit("bye", last)
		server.BroadcastToRoom("/", "chat1", "notice", "广播通知")
		s.Close()
		return last
	})

	router.GET("/bcast", func(context *gin.Context) {
		// 向房间内的所有人员发消息
		server.BroadcastToRoom("/", "chat1", "notice", "广播通知")
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatal("listen serve: ", err)
		}
	}()
	//defer server.Close()
	router.Use(gin.Recovery())
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/public", http.Dir("../asset"))

}

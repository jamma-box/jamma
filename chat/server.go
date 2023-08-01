package chat

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func Register(router *gin.RouterGroup) {

	server := socketio.NewServer(nil)
	// redis 适配器
	//ok, err := server.Adapter(&socketio.RedisAdapterOptions{
	//	Addr:    "127.0.0.1:6379",
	//	Prefix:  "socket.io",
	//	Network: "tcp",
	//})
	//
	//log.Println("redis:", ok)
	//
	//if err != nil {
	//	log.Fatal("error:", err)
	//}
	// 建立连接
	server.OnConnect("/", func(s socketio.Conn) error {
		//params, _ := url.ParseQuery(s.URL().RawQuery)
		//uid := params.Get("id")
		//s.SetContext(uid)
		////加入房间
		//s.Join("chat1")
		log.Println("建立连接::", s.ID())
		s.Emit("notice", "成功")
		return nil
	})
	// 连接错误
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("连接错误:", s.ID(), e) //记录连接错误信息
	})
	// 断开连接
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		s.LeaveAll() //将socket从所有加入的room中移除
		if uid := s.Context(); uid != nil {
			log.Printf("用户[%s]断开连接", uid)
		}
		log.Println("关闭连接：", s.ID(), reason)
	})

	// 服务/事件
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg) // 向client回复内容
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

	defer server.Close()
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/public", http.Dir("../asset"))

}

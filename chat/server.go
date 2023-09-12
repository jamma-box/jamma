package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

var server *socketio.Server

func Broadcast() {
	server.BroadcastToNamespace("/", "d", "")
}

func Register(router *gin.RouterGroup) {

	server = socketio.NewServer(nil)

	// 默认event
	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("[chat]OnConnect", s.ID())
		s.Join("public")
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("[chat]OnError", s.ID(), e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("[chat]OnDisconnect", s.ID(), reason)
	})

	// 服务/事件
	//notice用于内部的广播
	server.OnEvent("/", "notice", func(s socketio.Conn, msg any) {
		s.Emit("notice", msg) // 向client回复内容
	})

	//离开房间
	server.OnEvent("/", "bye", func(s socketio.Conn, msg Text) {
		s.Leave(msg.Room)
		//向订阅发布离开消息
		//s.Emit("bye", fmt.Sprintf("用户%v已离开%v房间", msg.Id, msg.Content))
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})

	////加入房间
	server.OnEvent("/", "join", func(s socketio.Conn, msg Text) {
		s.Join(msg.Room)
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})

	//文本消息
	server.OnEvent("/", "text", func(s socketio.Conn, msg Text) {

		//消息广播
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})

	//红包消息
	server.OnEvent("/", "red-packet", func(s socketio.Conn, msg RedPacket) {
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})

	//抢红包
	server.OnEvent("/", "grab-packet", func(s socketio.Conn, msg GrabPacket) {
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatal("listen serve: ", err)
		}
	}()

	//websocket服务
	router.Any("/socket.io/*any", gin.WrapH(server))

	//查看历史
	router.GET("/history/:room", func(ctx *gin.Context) {

	})

}

package chat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"jamma/types"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               // 请求方法
		origin := c.Request.Header.Get("Origin") // 请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)                                    // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") // 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                   //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

func Register(router *gin.RouterGroup) {

	server := socketio.NewServer(nil)
	// 默认event
	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("连接成功：")
		s.Join("chat")
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("连接错误：", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("关闭连接：", reason)
	})

	// 服务/事件
	//notice用于内部的广播
	server.OnEvent("/", "notice", func(s socketio.Conn, msg any) {
		s.Emit("notice", msg) // 向client回复内容
	})

	//离开房间
	server.OnEvent("/", "bye", func(s socketio.Conn, msg types.ChatText) {
		s.Leave(msg.Room)
		//向订阅发布离开消息
		//s.Emit("bye", fmt.Sprintf("用户%v已离开%v房间", msg.Id, msg.Content))
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})

	////加入房间
	server.OnEvent("/", "join", func(s socketio.Conn, msg types.ChatText) {
		s.Join(msg.Room)
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})
	//文本消息
	server.OnEvent("/", "text", func(s socketio.Conn, msg types.ChatText) {
		//添加到记录中
		err := textRecord(msg)
		if err != nil {
			s.Emit("notice", err.Error())
			return
		}
		//消息广播
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})
	//图片消息
	server.OnEvent("/", "img", func(s socketio.Conn, msg types.ChatImg) {
		////D:\\GoCode\\jamma/static/2023/8/41127461
		//file, err := os.ReadFile(msg.ImgPath)
		//if err != nil {
		//	s.Emit("notice", errors.New("图片获取失败"))
		//	return
		//}
		//添加到记录中
		err := imgRecord(msg)
		if err != nil {
			s.Emit("notice", err.Error())
			return
		}
		//图片的广播
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})
	//红包消息
	server.OnEvent("/", "hb", func(s socketio.Conn, msg types.ChatHongBao) {
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})
	//抢红包
	server.OnEvent("/", "qhb", func(s socketio.Conn, msg types.ChatQiangHongBao) {
		server.BroadcastToRoom("/", msg.Room, "notice", msg)
	})
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatal("listen serve: ", err)
		}
	}()
	router.Use(Cors())
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	//router.StaticFS("/public", http.Dir("../asset"))

}
func textRecord(msg types.ChatText) error {
	//路径
	filename := "/static/record"
	getwd, err := os.Getwd()
	t := time.Now()
	if err != nil {
		return errors.New("根目录读取失败")
	}
	path := filepath.Join(getwd, filename, fmt.Sprintf("%v", t.Year()), fmt.Sprintf("%v", int(t.Month())))

	//创建目录
	_ = os.MkdirAll(path, 777)
	//读写文件
	file, err := os.OpenFile(path+"/record.txt", os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		return err
	}
	defer file.Close()
	//追加内容

	//转换成json字符串
	data, err := json.Marshal(types.TextRecord{
		ChatText: types.ChatText{
			Id:      msg.Id,
			Room:    msg.Room,
			Content: msg.Content,
		},
		Type: "text",
		Time: t.Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return errors.New("数据编码错误")
	}
	//写入
	_, err = file.Write(data)
	if err != nil {
		return errors.New("文件写入失败")
	}
	//写入消息列表中
	return nil
}
func imgRecord(msg types.ChatImg) error {
	//路径
	filename := "/static/record"
	getwd, err := os.Getwd()
	t := time.Now()
	if err != nil {
		return errors.New("根目录读取失败")
	}
	path := filepath.Join(getwd, filename, fmt.Sprintf("%v", t.Year()), fmt.Sprintf("%v", int(t.Month())))
	//创建目录
	_ = os.MkdirAll(path, 777)
	//读写文件
	file, err := os.OpenFile(path+"/record.txt", os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		return err
	}
	defer file.Close()

	//转换成json字符串
	data, err := json.Marshal(types.ImgRecord{
		Id:      msg.Id,
		Room:    msg.Room,
		ImgPath: msg.ImgPath,
		Type:    "image",
		Time:    t.Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return errors.New("数据编码错误")
	}
	//写入
	_, err = file.Write(data)
	if err != nil {
		return errors.New("文件写入失败")
	}
	//写入消息列表中
	return nil
}

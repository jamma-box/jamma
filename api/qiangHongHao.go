package api

import (
	"arcade/chat"
	"arcade/types"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"sync"
)

// 抢红包参数
type QiangHongBao struct {
	Id     string `json:"id"` //红包id
	UserId string `json:"user_id"`
}

var mutex = sync.Mutex{}

func qiangHongbaoRouter(app *gin.RouterGroup) {

	app.POST("/create", func(c *gin.Context) {
		qhb := new(QiangHongBao)
		err := c.ShouldBindJSON(qhb)
		if err != nil {
			curd.Error(c, errors.New("参数获取失败"))
			return
		}
		//	go携程 ：处理
		go qiangHongBao(c, qhb)

	})

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[chat.RedPacket]())

	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[chat.RedPacket](nil, nil,
		"id", "user_id", "money", "room"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[chat.RedPacket](nil, nil))

}
func qiangHongBao(c *gin.Context, qhb *QiangHongBao) {
	// 判断锁：
	// 		未上锁: 加锁 defer 解锁
	//      已上锁：等待3秒，每隔50ms判断一次锁，超时则提示稍后重试
	mutex.Lock()
	//数据获取
	//获取红包数据
	hb := new(chat.RedPacket)
	_, err := db.Engine.ID(qhb.Id).Get(hb)
	if err != nil {
		curd.Error(c, errors.New("红包数据获取错误"))
		return
	}
	//获取用户数据
	user := new(types.User)
	_, err = db.Engine.ID(qhb.UserId).Cols("balance").Get(user)
	if err != nil {
		curd.Error(c, errors.New("用户数据获取错误"))
		return
	}
	//	根据算法生成红包金额
	var money float64
	//if hb.Type == 0 {
	//	//	默认红包
	//	money = hb.TotalMoney / hb.TotalNum)
	//
	//} else if hb.Type == 1 {
	//	//	根据算法生成红包金额
	//	money = (hb.TotalMoney / hb.TotalNum) * rand.Float64()
	//}
	////加钱扣钱
	//hb.CurrentMoney -= money
	hb.CurrentNum--
	//user.Balance += money
	//抢红包记录
	qianghongbao := chat.GrabPacket{
		Id:     qhb.Id,
		UserId: qhb.UserId,
		Money:  money,
		Room:   hb.Room,
	}
	_, err2 := db.Engine.ID(qhb.Id).Insert(&qianghongbao)
	if err2 != nil {
		curd.Error(c, errors.New("抢红包数据记录失败"))
		return
	}
	//红包数据更新
	_, err2 = db.Engine.ID(qhb.Id).Update(hb)
	if err2 != nil {
		curd.Error(c, errors.New("红包数据更新失败"))
		return
	}
	//用户数据更新
	_, err2 = db.Engine.ID(qhb.UserId).Update(user)
	if err2 != nil {
		curd.Error(c, errors.New("用户数据更新失败"))
		return
	}
	//返回抢红包红包状态
	curd.OK(c, qianghongbao)
	mutex.Unlock()

}

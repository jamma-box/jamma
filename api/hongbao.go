package api

import (
	"arcade/chat"
	"arcade/types"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"math/rand"
	"time"
)

// 红包请求参数
type HongBaoReq struct {
	UserId      string  `json:"user_id"`
	Money       float64 `json:"money"`    //红包金额
	Num         int64   `json:"num"`      //红包数量
	PayPassword string  `json:"password"` //支付密码
	Type        int64   `json:"type"`     //红包类型,0:默认红包，1：随机红包
	Room        string  `json:"room"`
}

func hongbaoRouter(app *gin.RouterGroup) {

	app.POST("/create", func(c *gin.Context) {
		//绑定参数
		hongbao := new(HongBaoReq)
		err := c.ShouldBindJSON(hongbao)
		if err != nil {
			curd.Error(c, err)
			return
		}
		//查询数据
		user := new(types.User)
		_, err = db.Engine.ID(hongbao.UserId).Get(&user)
		if err != nil {
			curd.Error(c, err)
			return
		}
		//if user.PayPassword == "" {
		//	curd.Error(c, errors.New("支付密码未设置，请去用户界面设置"))
		//	return
		//} else if user.PayPassword != hongbao.PayPassword {
		//	curd.Error(c, errors.New("支付密码错误"))
		//	return
		//}

		//	扣钱
		user.Balance -= hongbao.Money
		_, err = db.Engine.ID(hongbao.UserId).Update(&user)
		if err != nil {
			curd.Error(c, err)
			return
		}
		//红包记录
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomNumber := generateRandomNumber(10, r)

		hb := chat.RedPacket{
			Id:           randomNumber,
			UserId:       hongbao.UserId,
			Type:         hongbao.Type,
			Room:         hongbao.Room,
			CurrentMoney: hongbao.Money,
			CurrentNum:   hongbao.Num,
			TotalMoney:   hongbao.Money,
			TotalNum:     hongbao.Num,
		}
		_, err = db.Engine.InsertOne(&hb)
		if err != nil {
			curd.Error(c, err)
			return
		}
		//返回红包数据
		curd.OK(c, hb)
	})

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[chat.RedPacket]())

	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[chat.RedPacket](nil, nil,
		"user_id", "room", "type", "current_money", "current_num", "total_money", "total_num"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[chat.RedPacket](nil, nil))
}
func generateRandomNumber(length int, r *rand.Rand) string {
	// 随机数生成的字符集合
	charset := "0123456789"

	randomNumber := make([]byte, length)
	randomNumber[0] = charset[r.Intn(len(charset)-1)+1]
	for i := 1; i < length; i++ {
		randomNumber[i] = charset[r.Intn(len(charset))]
	}

	return string(randomNumber)
}

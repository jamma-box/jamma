package api

import (
	"arcade/types"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

// @Summary 查询消费记录数量
// @Schemes
// @Description 查询消费记录数量
// @Tags exchange
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回消费记录数量
// @Router /exchange/count [post]
func noopExchangeCount() {}

// @Summary 查询消费记录
// @Schemes
// @Description 查询消费记录
// @Tags exchange
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Exchange] 返回消费记录信息
// @Router /exchange/search [post]
func noopExchangeSearch() {}

// @Summary 查询消费记录
// @Schemes
// @Description 查询消费记录
// @Tags exchange
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Exchange] 返回消费记录信息
// @Router /exchange/list [get]
func noopExchangeList() {}

// @Summary 创建消费记录
// @Schemes
// @Description 创建消费记录
// @Tags exchange
// @Param search body types.Exchange true "消费记录信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Exchange] 返回消费记录信息
// @Router /exchange/create [post]
func noopExchangeCreate() {}

// @Summary 获取消费记录
// @Schemes
// @Description 获取消费记录
// @Tags exchange
// @Param id path int true "消费记录ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Exchange] 返回消费记录信息
// @Router /exchange/{id} [get]

func noopExchangeGet() {}
func exchangeRouter(app *gin.RouterGroup) {
	app.POST("/count", curd.ApiCount[types.Exchange]())

	app.POST("/search", curd.ApiSearch[types.Exchange]())

	app.GET("/list", curd.ApiList[types.Exchange]())

	app.POST("/create", curd.ApiCreateHook[types.Exchange](func(m *types.Exchange) error {
		var user types.User
		has, err := db.Engine.ID(m.UserId).Get(&user)
		if err != nil {
			return err
		}
		if !has {
			return errors.New("无此用户")
		}
		if user.Balance < m.Amount {
			return errors.New("余额不足")
		}
		user.Balance = user.Balance - m.Amount
		_, err = db.Engine.Cols("balance").Update(&user)
		return err
	}, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.Exchange]())
}

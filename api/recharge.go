package api

import (
	"arcade/types"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询重置记录数量
// @Schemes
// @Description 查询重置记录数量
// @Tags recharge
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回重置记录数量
// @Router /recharge/count [post]
func noopRechargeCount() {}

// @Summary 查询重置记录
// @Schemes
// @Description 查询重置记录
// @Tags recharge
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Recharge] 返回重置记录信息
// @Router /recharge/search [post]
func noopRechargeSearch() {}

// @Summary 查询重置记录
// @Schemes
// @Description 查询重置记录
// @Tags recharge
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Recharge] 返回重置记录信息
// @Router /recharge/list [get]
func noopRechargeList() {}

// @Summary 创建重置记录
// @Schemes
// @Description 创建重置记录
// @Tags recharge
// @Param search body types.Recharge true "重置记录信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Recharge] 返回重置记录信息
// @Router /recharge/create [post]
func noopRechargeCreate() {}

// @Summary 获取重置记录
// @Schemes
// @Description 获取重置记录
// @Tags recharge
// @Param id path int true "重置记录ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Recharge] 返回重置记录信息
// @Router /recharge/{id} [get]

func noopRechargeGet() {}
func rechargeRouter(app *gin.RouterGroup) {
	app.POST("/count", curd.ApiCount[types.Recharge]())

	app.POST("/search", curd.ApiSearch[types.Recharge]())

	app.GET("/list", curd.ApiList[types.Recharge]())

	app.POST("/create", curd.ApiCreateHook[types.Recharge](nil, nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Recharge]())
}

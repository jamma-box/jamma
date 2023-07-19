package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"jamma/types"
)

// @Summary 查询签到记录数量
// @Schemes
// @Description 查询签到记录数量
// @Tags device
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回签到记录数量
// @Router /record/signin/count [post]
func noopSignInCount() {}

// @Summary 查询签到记录
// @Schemes
// @Description 查询签到记录
// @Tags device
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.SignIn] 返回签到记录信息
// @Router /record/signin/search [post]
func noopSignInSearch() {}

// @Summary 查询签到记录
// @Schemes
// @Description 查询签到记录
// @Tags device
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.SignIn] 返回签到记录信息
// @Router /record/signin/list [get]
func noopSignInList() {}

// @Summary 创建签到记录
// @Schemes
// @Description 创建签到记录
// @Tags device
// @Param search body types.SignIn true "签到记录信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.SignIn] 返回签到记录信息
// @Router /record/signin/create [post]
func noopSignInCreate() {}

// @Summary 获取签到记录
// @Schemes
// @Description 获取签到记录
// @Tags device
// @Param id path int true "签到记录ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.SignIn] 返回签到记录信息
// @Router /record/signin/{id} [get]
func noopSignInGet() {}

func signRouter(app *gin.RouterGroup) {
	app.POST("/count", curd.ApiCount[types.SignIn]())

	app.POST("/search", curd.ApiSearch[types.Recharge]())

	app.GET("/list", curd.ApiList[types.SignIn]())

	app.POST("/create", curd.ApiCreateHook[types.SignIn](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.SignIn]())
}

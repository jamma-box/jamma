package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"jamma/types"
)

// @Summary 查询游戏机数量
// @Schemes
// @Description 查询游戏机数量
// @Tags device
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回游戏机数量
// @Router /device/count [post]
func noopDeviceCount() {}

// @Summary 查询游戏机
// @Schemes
// @Description 查询游戏机
// @Tags device
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.GameBox] 返回游戏机信息
// @Router /device/search [post]
func noopDeviceSearch() {}

// @Summary 查询游戏机
// @Schemes
// @Description 查询游戏机
// @Tags device
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.GameBox] 返回游戏机信息
// @Router /device/list [get]
func noopDeviceList() {}

// @Summary 创建游戏机
// @Schemes
// @Description 创建游戏机
// @Tags device
// @Param search body types.GameBox true "游戏机信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.GameBox] 返回游戏机信息
// @Router /device/create [post]
func noopDeviceCreate() {}

// @Summary 获取游戏机
// @Schemes
// @Description 获取游戏机
// @Tags device
// @Param id path int true "游戏机ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.GameBox] 返回游戏机信息
// @Router /device/{id} [get]
func noopDeviceGet() {}

// @Summary 修改游戏机
// @Schemes
// @Description 修改游戏机
// @Tags device
// @Param id path int true "游戏机ID"
// @Param device body types.GameBox true "游戏机信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.GameBox] 返回游戏机信息
// @Router /device/{id} [post]
func noopDeviceUpdate() {}

// @Summary 删除游戏机
// @Schemes
// @Description 删除游戏机
// @Tags device
// @Param id path int true "游戏机ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.GameBox] 返回游戏机信息
// @Router /device/{id}/delete [get]
func noopDeviceDelete() {}


func deviceRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.GameBox]())

	app.POST("/search", curd.ApiSearch[types.GameBox]("id", "name", "disabled", "created"))

	app.GET("/list", curd.ApiList[types.GameBox]())

	app.POST("/create", curd.ApiCreateHook[types.GameBox](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.GameBox]())

	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[types.GameBox](nil, nil,
		 "type", "name", "desc", "disabled"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.GameBox](nil, nil))

	app.GET(":id/disable", curd.ParseParamId, curd.ApiDisableHook[types.GameBox](true, nil, nil))

	app.GET(":id/enable", curd.ParseParamId, curd.ApiDisableHook[types.GameBox](false, nil, nil))
}

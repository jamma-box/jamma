package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"jamma/types"
)

// @Summary 查询游戏厅数量
// @Schemes
// @Description 查询游戏厅数量
// @Tags device
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回游戏厅数量
// @Router /hall/count [post]
func noopGameHallCount() {}

// @Summary 查询游戏厅
// @Schemes
// @Description 查询游戏厅
// @Tags device
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.GameHall] 返回游戏厅信息
// @Router /hall/search [post]
func noopGameHallSearch() {}

// @Summary 查询游戏厅
// @Schemes
// @Description 查询游戏厅
// @Tags device
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.GameHall] 返回游戏厅信息
// @Router /hall/list [get]
func noopGameHallList() {}

// @Summary 创建游戏厅
// @Schemes
// @Description 创建游戏厅
// @Tags device
// @Param search body types.GameHall true "游戏厅信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.GameHall] 返回游戏厅信息
// @Router /hall/create [post]
func noopGameHallCreate() {}

// @Summary 获取游戏厅
// @Schemes
// @Description 获取游戏厅
// @Tags device
// @Param id path int true "游戏厅ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.GameHall] 返回游戏厅信息
// @Router /hall/{id} [get]
func noopGameHallGet() {}

// @Summary 修改游戏厅
// @Schemes
// @Description 修改游戏厅
// @Tags device
// @Param id path int true "游戏厅ID"
// @Param device body types.GameHall true "游戏厅信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.GameHall] 返回游戏厅信息
// @Router /hall/{id} [post]
func noopGameHallUpdate() {}

// @Summary 删除游戏厅
// @Schemes
// @Description 删除游戏厅
// @Tags device
// @Param id path int true "游戏厅ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.GameHall] 返回游戏厅信息
// @Router /hall/{id}/delete [get]
func noopGameHallDelete() {}

func gamehallRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.GameHall]())

	app.POST("/search", curd.ApiSearch[types.GameHall]("id", "name", "disabled", "created"))

	app.GET("/list", curd.ApiList[types.GameHall]())

	app.POST("/create", curd.ApiCreateHook[types.GameHall](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.GameHall]())

	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[types.GameHall](nil, nil,
		"name", "desc", "img", "disabled"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.GameHall](nil, nil))

	app.GET(":id/disable", curd.ParseParamId, curd.ApiDisableHook[types.GameHall](true, nil, nil))

	app.GET(":id/enable", curd.ParseParamId, curd.ApiDisableHook[types.GameHall](false, nil, nil))
}

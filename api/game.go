package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"jamma/types"
)

// @Summary 查询游戏厅数量
// @Schemes
// @Description 查询游戏厅数量
// @Tags game
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回游戏厅数量
// @Router /hall/count [post]
func noopGameCount() {}

// @Summary 查询游戏厅
// @Schemes
// @Description 查询游戏厅
// @Tags game
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Game] 返回游戏厅信息
// @Router /hall/search [post]
func noopGameSearch() {}

// @Summary 查询游戏厅
// @Schemes
// @Description 查询游戏厅
// @Tags game
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Game] 返回游戏厅信息
// @Router /hall/list [get]
func noopGameList() {}

// @Summary 创建游戏厅
// @Schemes
// @Description 创建游戏厅
// @Tags game
// @Param search body types.Game true "游戏厅信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Game] 返回游戏厅信息
// @Router /hall/create [post]
func noopGameCreate() {}

// @Summary 获取游戏厅
// @Schemes
// @Description 获取游戏厅
// @Tags game
// @Param id path int true "游戏厅ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Game] 返回游戏厅信息
// @Router /hall/{id} [get]
func noopGameGet() {}

// @Summary 修改游戏厅
// @Schemes
// @Description 修改游戏厅
// @Tags game
// @Param id path int true "游戏厅ID"
// @Param game body types.Game true "游戏厅信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Game] 返回游戏厅信息
// @Router /hall/{id} [post]
func noopGameUpdate() {}

// @Summary 删除游戏厅
// @Schemes
// @Description 删除游戏厅
// @Tags game
// @Param id path int true "游戏厅ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Game] 返回游戏厅信息
// @Router /hall/{id}/delete [get]
func noopGameDelete() {}

func gameRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Game]())

	app.POST("/search", curd.ApiSearch[types.Game]())

	app.GET("/list", curd.ApiList[types.Game]())

	app.POST("/create", curd.ApiCreateHook[types.Game](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.Game]())

	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[types.Game](nil, nil,
		"name", "desc", "icon", "type", "disabled"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.Game](nil, nil))

	app.GET(":id/disable", curd.ParseParamId, curd.ApiDisableHook[types.Game](true, nil, nil))

	app.GET(":id/enable", curd.ParseParamId, curd.ApiDisableHook[types.Game](false, nil, nil))
}

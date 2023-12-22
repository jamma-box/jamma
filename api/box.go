package api

import (
	"arcade/box"
	"arcade/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"strconv"
)

// @Summary 查询游戏机数量
// @Schemes
// @Description 查询游戏机数量
// @Tags box
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回游戏机数量
// @Router /box/count [post]
func noopBoxCount() {}

// @Summary 查询游戏机
// @Schemes
// @Description 查询游戏机
// @Tags box
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Box] 返回游戏机信息
// @Router /box/search [post]
func noopBoxSearch() {}

// @Summary 查询游戏机
// @Schemes
// @Description 查询游戏机
// @Tags box
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Box] 返回游戏机信息
// @Router /box/list [get]
func noopBoxList() {}

// @Summary 创建游戏机
// @Schemes
// @Description 创建游戏机
// @Tags box
// @Param search body types.Box true "游戏机信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Box] 返回游戏机信息
// @Router /box/create [post]
func noopBoxCreate() {}

// @Summary 获取游戏机
// @Schemes
// @Description 获取游戏机
// @Tags box
// @Param id path int true "游戏机ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Box] 返回游戏机信息
// @Router /box/{id} [get]
func noopBoxGet() {}

// @Summary 修改游戏机
// @Schemes
// @Description 修改游戏机
// @Tags box
// @Param id path int true "游戏机ID"
// @Param box body types.Box true "游戏机信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Box] 返回游戏机信息
// @Router /box/{id} [post]
func noopBoxUpdate() {}

// @Summary 删除游戏机
// @Schemes
// @Description 删除游戏机
// @Tags box
// @Param id path int true "游戏机ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Box] 返回游戏机信息
// @Router /box/{id}/delete [get]
func noopBoxDelete() {}

func boxRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Box]())

	app.POST("/search", curd.ApiSearch[types.Box]())

	app.GET("/list", curd.ApiList[types.Box]())

	app.POST("/create", curd.ApiCreateHook[types.Box](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.Box]())

	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[types.Box](nil, nil,
		"name", "desc", "icon", "type", "disabled", "game_id"))

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.Box](nil, nil))

	app.GET("/:id/disable", curd.ParseParamId, curd.ApiDisableHook[types.Box](true, nil, nil))

	app.GET("/:id/enable", curd.ParseParamId, curd.ApiDisableHook[types.Box](false, nil, nil))

	app.GET("/:id/seat/:seat", curd.ParseParamId, func(ctx *gin.Context) {
		b := box.Get(ctx.Param("id"))
		if b == nil {
			curd.Fail(ctx, "找不到设备")
			return
		}

		seat, err := strconv.Atoi(ctx.Param("seat"))
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		if b.Seats[seat].UserId != 0 {
			//TODO 判断位置重复
			curd.Fail(ctx, "已占位")
			return
		}

		//坐下
		b.Seats[seat].UserId = ctx.GetInt64("user")

		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		defer c.Close()

		b.Seat(c, seat)
	})

	app.GET("/:id/bridge", curd.ParseParamId, func(ctx *gin.Context) {
		b := box.Get(ctx.Param("id"))
		if b == nil {
			curd.Fail(ctx, "找不到设备")
			return
		}

		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		defer c.Close()

		b.Bridge(c)
	})

	app.GET("/:id/live", curd.ParseParamId, func(ctx *gin.Context) {
		b := box.Get(ctx.Param("id"))
		if b == nil {
			curd.Fail(ctx, "找不到设备")
			return
		}

		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		defer c.Close()

		b.Live(c)
	})

	app.GET("/:id/pad", curd.ParseParamId, func(ctx *gin.Context) {
		b := box.Get(ctx.Param("id"))
		if b == nil {
			curd.Fail(ctx, "找不到设备")
			return
		}

		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		defer c.Close()

		b.Pad(c)
	})
}

var upgrader = websocket.Upgrader{} // use default options

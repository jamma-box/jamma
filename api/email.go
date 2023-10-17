package api

import (
	"arcade/types"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

func emailRouter(app *gin.RouterGroup) {
	app.POST("/count", curd.ApiCount[types.Email]())

	app.POST("/search", curd.ApiSearch[types.Email]())

	app.GET("/list", curd.ApiList[types.Email]())

	app.POST("/create", curd.ApiCreateHook[types.Email](nil, nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Email]())

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Email](nil, nil))

	app.GET("/:id/read", emailRead)

}

func emailRead(ctx *gin.Context) {
	//id := ctx.GetInt64("id")
	id := ctx.MustGet("id")
	var data types.Email
	data.Read = true
	_, err := db.Engine.ID(id).Cols("read").Update(&data)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, nil)
}

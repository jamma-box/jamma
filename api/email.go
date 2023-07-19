package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"jamma/types"
	"reflect"
)

func emailRouter(app *gin.RouterGroup) {
	app.POST("/count", curd.ApiCount[types.Email]())

	app.POST("/search", curd.ApiSearch[types.Email]())

	app.GET("/list", curd.ApiList[types.Email]())

	app.POST("/create", curd.ApiCreateHook[types.Email](nil, nil))

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.Email]())

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.Email](nil, nil))

	app.GET("/:id/read", curd.ParseParamId, readHook[types.Email](true, nil, nil))
	app.GET("/:id/unread", curd.ParseParamId, readHook[types.Email](false, nil, nil))

}
func readHook[T any](read bool, before, after func(id any) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id := ctx.GetInt64("id")
		id := ctx.MustGet("id")
		if before != nil {
			if err := before(id); err != nil {
				curd.Error(ctx, err)
				return
			}
		}
		var data T
		value := reflect.ValueOf(&data).Elem()
		field := value.FieldByName("Read")
		field.SetBool(read)

		_, err := db.Engine.ID(id).Cols("read").Update(&data)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		if after != nil {
			if err := after(id); err != nil {
				curd.Error(ctx, err)
				return
			}
		}

		curd.OK(ctx, nil)
	}
}

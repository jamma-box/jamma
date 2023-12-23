package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"net/http"
)

func CatchError(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//runtime.Stack()
			//debug.Stack()
			switch err.(type) {
			case error:
				curd.Error(ctx, err.(error))
			case string:
				curd.Fail(ctx, err.(string))
			default:
				ctx.JSON(http.StatusOK, gin.H{"error": err})
			}
		}
	}()
	ctx.Next()
}

func MustLogin(ctx *gin.Context) {
	token := ctx.Request.URL.Query().Get("token")
	if token == "" {
		token = ctx.Request.Header.Get("Authorization")
		if token != "" {
			//此处需要去掉 Bearer
			token = token[7:]
		}
	}

	if token != "" {
		claims, err := JwtVerify(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
		} else {
			ctx.Set("user", claims.Id) //与session统一
			ctx.Next()
		}
		return
	}

	//if claims.ExpiresAt < time.Now().Unix() {
	//	curd.Error(c, fmt.Errorf("令牌已过期"))
	//	c.Abort()
	//	return
	//}

	curd.Fail(ctx, "令牌验证失败")
	ctx.Abort()
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func RegisterRoutes(router *gin.RouterGroup) {

	//错误恢复，并返回至前端

	router.POST("/login", login)
	router.GET("/weixin/auth", weixinAuth)

	//游戏机
	router.GET("/box/:id/live", curd.ParseParamStringId, boxLive)
	router.GET("/box/:id/pad", curd.ParseParamStringId, boxPad)

	//检查 session，必须登录
	router.Use(MustLogin)

	router.GET("/logout", logout)

	router.POST("/password", password)

	weixinRouter(router.Group("/weixin"))

	//注册子接口
	userRouter(router.Group("/user"))

	backupRouter(router.Group("/backup"))

	attachRouter(router.Group("/attach"))

	boxRouter(router.Group("/box"))

	gameRouter(router.Group("/game"))

	rechargeRouter(router.Group("/recharge"))

	signRouter(router.Group("/sign"))

	imgRouter(router.Group("/img"))

	emailRouter(router.Group("/email"))

	hongbaoRouter(router.Group("/hongbao"))

	qiangHongbaoRouter(router.Group("/hongbao/qiang"))

	router.Use(func(ctx *gin.Context) {
		curd.Fail(ctx, "Not found")
		ctx.Abort()
	})
}

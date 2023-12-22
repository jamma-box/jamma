package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"net/http"
	"strings"
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

func MustLogin(c *gin.Context) {
	// 检查有没有Bearer前缀
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		tokenString = c.GetHeader("authorization")
	}
	if tokenString == "" {
		curd.Error(c, fmt.Errorf("UnAuthorized"))
		c.Abort()
		return
	}
	if !strings.HasPrefix(tokenString, "Bearer ") {
		curd.Error(c, fmt.Errorf("令牌错误格式"))
		c.Abort()
		return
	}
	// 解析
	claims, err := JwtVerify(tokenString[7:])
	if err != nil {
		curd.Error(c, fmt.Errorf("UnAuthorized err:%v", err.Error()))
		c.Abort()
		return
	}

	//if claims.ExpiresAt < time.Now().Unix() {
	//	curd.Error(c, fmt.Errorf("令牌已过期"))
	//	c.Abort()
	//	return
	//}

	if claims.Id == 0 {
		session := sessions.Default(c)
		if id := session.Get("user"); id == nil {
			curd.Error(c, fmt.Errorf("令牌验证失败"))
			c.Abort()
			return
		} else {
			c.Set("user", id)
			c.Next()
		}
	} else {
		c.Set("user", claims.Id)
		c.Next()
	}

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

package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"net/http"
)

func catchError(ctx *gin.Context) {
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
			//TODO 这里好像又继续了
		}
	}()
	ctx.Next()

	//TODO 内容如果为空，返回404

}

func mustLogin(ctx *gin.Context) {
	token := ctx.Request.URL.Query().Get("token")
	if token == "" {
		token = ctx.Request.Header.Get("Authorization")
		if token != "" {
			//此处需要去掉 Bearer
			token = token[7:]
		}
	}

	if token != "" {
		claims, err := jwtVerify(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
		}
		ctx.Set("user", claims.Id) //与session统一
		ctx.Next()
		return
	}

	//检查Session
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Set("user", user)
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
	}
}

func RegisterRoutes(router *gin.RouterGroup) {
	//错误恢复，并返回至前端
	router.Use(catchError)

	router.GET("/auth", auth)
	router.POST("/login", login)

	//检查 session，必须登录
	router.Use(mustLogin)

	router.GET("/logout", logout)
	router.POST("/password", password)


	//注册子接口
	userRouter(router.Group("/user"))

	deviceRouter(router.Group("/device"))

	backupRouter(router.Group("/backup"))

	attachRouter(router.Group("/attach"))

	//TODO 报接口错误（以下代码不生效，路由好像不是树形处理）
	router.Use(func(ctx *gin.Context) {
		curd.Fail(ctx, "Not found")
		ctx.Abort()
	})
}

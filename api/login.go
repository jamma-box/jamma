package api

import (
	"arcade/types"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

type loginObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func md5hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}

func login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var obj loginObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {

		curd.Error(ctx, err)
		return
	}
	var user types.User
	has, err := db.Engine.Where("username=?", obj.Username).Get(&user)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	if !has {
		//管理员自动创建
		if obj.Username == "admin" {
			user.Id = "0"
			user.Username = obj.Username
			user.Nickname = "管理员"
			_, err = db.Engine.InsertOne(&user)
			if err != nil {
				curd.Error(ctx, err)
				return
			}
		} else {
			curd.Fail(ctx, "找不到用户")
			return
		}
	}

	if user.Disabled {
		curd.Fail(ctx, "用户已禁用")
		return
	}

	var password types.Password
	has, err = db.Engine.ID(user.Id).Get(&password)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//初始化密码
	if !has {
		dp := "123456"

		password.Id = user.Id
		password.Password = md5hash(dp)
		_, err = db.Engine.InsertOne(&password)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
	}

	if password.Password != obj.Password {
		curd.Fail(ctx, "密码错误")
		return
	}

	//_, _ = db.Engine.InsertOne(&types.UserEvent{UserId: user.Id, ModEvent: types.ModEvent{Type: "登录"}})

	//存入session
	session.Set("user", user.Id)
	_ = session.Save()
	token, err := jwtGenerate(user.Id)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	res := make(map[string]interface{}, 2)
	res["user"] = user
	res["token"] = token
	curd.OK(ctx, res)
}

func logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	u := session.Get("user")
	if u == nil {
		curd.Fail(ctx, "未登录")
		return
	}

	//user := u.(int64)
	//_, _ = db.Engine.InsertOne(&types.UserEvent{UserId: user, ModEvent: types.ModEvent{Type: "退出"}})

	session.Clear()
	_ = session.Save()
	curd.OK(ctx, nil)
}

type passwordObj struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func password(ctx *gin.Context) {

	var obj passwordObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		curd.Error(ctx, err)
		return
	}
	if obj.Old == "" {
		curd.Fail(ctx, "密码错误")
		return
	}
	if obj.New == "" {
		curd.Fail(ctx, "新密码不能为空")
		return
	}
	var pwd types.Password
	has, err := db.Engine.ID(ctx.GetString("user")).Get(&pwd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	if !has {
		curd.Fail(ctx, "用户不存在")
		return
	}
	if obj.Old != pwd.Password {
		curd.Fail(ctx, "密码错误")
		return
	}

	pwd.Password = obj.New //前端已经加密过
	_, err = db.Engine.ID(ctx.GetString("user")).Cols("password").Update(&pwd)
	if err != nil {
		curd.Error(ctx, fmt.Errorf("密码修改失败，err:%v", err.Error()))
		return
	}

	curd.OK(ctx, nil)
}

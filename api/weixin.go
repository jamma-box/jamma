package api

import (
	"arcade/types"
	"arcade/weixin"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/pay/order"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"log"
	"strconv"
)

func weixinAuthRouter(app *gin.RouterGroup) {

	app.GET("/auth", weixinAuth)
}

func weixinRouter(app *gin.RouterGroup) {
	app.GET("/pre-pay", weixinPrePay)
	app.GET("/get-js", weixinGetJS)
}

func weixinGetJS(ctx *gin.Context) {
	js := weixin.GetOfficialAccount().GetJs()
	config, err := js.GetConfig(ctx.Query("url"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, config)
}

func weixinPrePay(ctx *gin.Context) {

	var u types.User
	_, err := db.Engine.ID(ctx.GetInt64("user")).Get(&u)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	amount, err := strconv.Atoi(ctx.Query("amount"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	o := types.Recharge{
		UserId: u.Id,
		Amount: int64(amount),
	}

	_, err = db.Engine.InsertOne(&o)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	log.Println("order", &o)

	ord, err := weixin.GetPay().GetOrder().PrePayOrder(&order.Params{
		TotalFee:   strconv.Itoa(amount),
		CreateIP:   "127.0.0.1",
		OutTradeNo: strconv.Itoa(int(o.Id)),
		OpenID:     u.OpenId,
		TradeType:  "JSAPI",
		Body:       "充值",
		//Detail:     "充值",
		NotifyURL: "https://gamebox.zgwit.cn/pay_notify",
	})
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, ord)
}

func weixinAuth(ctx *gin.Context) {
	code := ctx.Query("code")

	oauth := weixin.GetOfficialAccount().GetOauth()
	token, err := oauth.GetUserAccessToken(code)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	var user types.User
	has, err := db.Engine.Where("openid=?", token.OpenID).Get(&user)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//自动创建用户
	if !has {
		info, err := oauth.GetUserInfo(token.AccessToken, token.OpenID, "zh_CN")
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		user.OpenId = info.OpenID
		user.Username = info.OpenID
		user.Nickname = info.Nickname
		user.Avatar = info.HeadImgURL
		user.Balance = 10 //送10个金币，可以从配置文件中取

		_, err = db.Engine.InsertOne(&user)
		if err != nil {
			curd.Error(ctx, err)
			return
		}
	}

	curd.OK(ctx, user)
}

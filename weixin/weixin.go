package weixin

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	oaConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/pay"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
)

var oa *officialaccount.OfficialAccount
var py *pay.Pay

func GetOfficialAccount() *officialaccount.OfficialAccount {
	return oa
}

func GetPay() *pay.Pay {
	return py
}

func Open() {

	wc := wechat.NewWechat()

	ca := cache.NewMemory()

	oa = wc.GetOfficialAccount(&oaConfig.Config{
		AppID:          "",
		AppSecret:      "",
		Token:          "",
		EncodingAESKey: "",
		Cache:          ca,
	})

	py = wc.GetPay(&payConfig.Config{
		AppID:     "",
		MchID:     "",
		Key:       "",
		NotifyURL: "",
	})
}

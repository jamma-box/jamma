package weixin

import (
	"github.com/zgwit/iot-master/v3/pkg/config"
)

// Options 参数
type Options struct {
	AppId     string     `json:"appid"`
	AppSecret string     `json:"app_secret"`
	Pay       PayOptions `json:"pay"`
}

type PayOptions struct {
	AppId     string `json:"appid"`
	MchId     string `json:"mch_id"`
	Key       string `json:"key"`
	NotifyUrl string `json:"notify_url"`
}

func Default() Options {
	return Options{}
}

var options Options = Default()
var configure = config.AppName() + ".weixin.yaml"

const ENV = "IOT_MASTER_WEIXIN_"

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}

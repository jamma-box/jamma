package config

import (
	"arcade/weixin"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

func Load() {
	err := log.Load()
	if err != nil {
		_ = log.Store()
	}
	err = web.Load()
	if err != nil {
		_ = web.Store()
	}
	err = db.Load()
	if err != nil {
		_ = db.Store()
	}
	err = weixin.Load()
	if err != nil {
		_ = weixin.Store()
	}
}

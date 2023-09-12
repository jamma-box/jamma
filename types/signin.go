package types

import (
	"time"
)

// 签到记录
type SignIn struct {
	Id      int64     `json:"id" xorm:"pk"`
	UserId  int64     `json:"user_id" xorm:"index"` //用户ID
	Created time.Time `json:"created" xorm:"created"`
}

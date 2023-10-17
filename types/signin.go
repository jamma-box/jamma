package types

import (
	"time"
)

// 签到记录
type SignIn struct {
	Id      string    `json:"id" xorm:"pk"`
	UserId  string    `json:"user_id" xorm:"index"` //用户ID
	Created time.Time `json:"created" xorm:"created"`
}

package types

import (
	"time"
)

// 签到记录
type SignIn struct {
	Id      int64     `json:"id,omitempty" form:"id"`
	UserId  string    `json:"user_id" form:"user_id"` //用户ID
	Created time.Time `json:"created" xorm:"created" form:"created"`
}

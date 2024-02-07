package types

import (
	"time"
)

// 消息记录
type Exchange struct {
	Id      int64     `json:"id,omitempty" form:"id"`
	UserId  int64     `json:"user_id" form:"user_id"` //用户ID
	Amount  int64     `json:"amount" form:"amount"`   //充值金额 分
	Type    string    `json:"type"`
	Phone   string    `json:"phone"`
	Result  bool      `json:"result"`
	Created time.Time `json:"created" xorm:"created" form:"created"`
}

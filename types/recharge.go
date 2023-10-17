package types

import (
	"time"
)

// 充值记录
type Recharge struct {
	Id      string    `json:"id" xorm:"pk"`
	UserId  string    `json:"user_id" xorm:"index"` //用户ID
	Amount  int64     `json:"amount"`               //充值金额 分
	Created time.Time `json:"created" xorm:"created"`
}

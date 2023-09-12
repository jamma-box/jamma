package types

import (
	"time"
)

// 充值记录
type Recharge struct {
	Id      int64     `json:"id" xorm:"pk"`
	UserId  int64     `json:"user_id" xorm:"index"` //用户ID
	Amount  int64     `json:"amount"`               //充值金额 分
	Created time.Time `json:"created" xorm:"created"`
}

package types

import (
	"time"
)

// 充值记录
type Recharge struct {
	Id      int64     `json:"id,omitempty" form:"id"`
	UserId  int64     `json:"user_id" form:"user_id"` //用户ID
	Amount  int64     `json:"amount" form:"amount"`   //充值金额 分
	Result  bool      `json:"result"`
	Created time.Time `json:"created" xorm:"created" form:"created"`
}

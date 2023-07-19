package types

import "time"

// 充值记录
type Recharge struct {
	Id        int64     `json:"id" xorm:"pk"`
	UserId    int64     `json:"user_id" xorm:"index"`   //用户ID
	Username  string    `json:"username" xorm:"index"`  //账号
	Cellphone string    `json:"cellphone" xorm:"index"` //手机
	Amount    int64     `json:"amount"`                 //充值金额
	Created   time.Time `json:"created" xorm:"created"`
}

// 签到记录
type SignIn struct {
	Id       int64     `json:"id" xorm:"pk"`
	UserId   int64     `json:"user_id" xorm:"index"`  //用户ID
	Username string    `json:"username" xorm:"index"` //账号
	Created  time.Time `json:"created" xorm:"created"`
}

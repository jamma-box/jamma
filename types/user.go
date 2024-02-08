package types

import "time"

// User 用户
type User struct {
	Id        int64     `json:"id,omitempty" form:"id"`
	OpenId    string    `json:"openid,omitempty" xorm:"openid" form:"openid"` //微信openid
	Username  string    `json:"username" xorm:"unique"  form:"username"`      //账号
	Nickname  string    `json:"nickname,omitempty" form:"nickname"`           //昵称
	Email     string    `json:"email,omitempty" form:"email"`                 //邮箱
	Cellphone string    `json:"cellphone,omitempty" form:"cellphone"`         //手机
	Avatar    string    `json:"avatar,omitempty" form:"avatar"`               //头像
	Balance   int64     `json:"balance" form:"balance"`                       //余额
	Gender    string    `json:"gender,omitempty" form:"gender"`               //性别
	Signature string    `json:"signature,omitempty" form:"signature"`         //签字
	Integral  int64     `json:"integral" form:"integral"`                     //游戏积分
	Disabled  bool      `json:"disabled" form:"disabled"`                     //禁用
	Created   time.Time `json:"created" xorm:"created" form:"created"`
}

type Me struct {
	User       `xorm:"extends" `
	Privileges []string `json:"privileges" form:"privileges"`
}

// Password 密码
type Password struct {
	Id       int64  `json:"id,omitempty" form:"id"`
	Password string `json:"password" form:"password"`
}

type UserHistory struct {
	Id      int64     `json:"id,omitempty"  form:"id"`
	UserId  string    `json:"user_id" form:"user_id"`
	Event   string    `json:"event" form:"event"`
	Created time.Time `json:"created" xorm:"created" form:"created"`
}

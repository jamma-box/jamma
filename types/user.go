package types

import "time"

// User 用户
type User struct {
	Id        string    `json:"id" xorm:"pk"`
	Username  string    `json:"username" xorm:"unique"` //账号
	Nickname  string    `json:"nickname,omitempty"`     //昵称
	Email     string    `json:"email,omitempty"`        //邮箱
	Cellphone string    `json:"cellphone,omitempty"`    //手机
	Avatar    string    `json:"avatar,omitempty"`       //头像
	Balance   float64   `json:"balance"`                //余额
	Integral  int64     `json:"integral"`               //游戏积分
	Disabled  bool      `json:"disabled"`               //禁用
	Created   time.Time `json:"created" xorm:"created"`
}

type Me struct {
	User       `xorm:"extends"`
	Privileges []string `json:"privileges"`
}

// Password 密码
type Password struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}

type UserHistory struct {
	Id      string    `json:"id"`
	UserId  string    `json:"user_id"`
	Event   string    `json:"event"`
	Created time.Time `json:"created" xorm:"created"`
}

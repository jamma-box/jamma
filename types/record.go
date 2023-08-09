package types

import "time"

// 充值记录
type Recharge struct {
	Id      int64     `json:"id" xorm:"pk"`
	UserId  int64     `json:"user_id" xorm:"index"` //用户ID
	Amount  int64     `json:"amount"`               //充值金额
	Created time.Time `json:"created" xorm:"created"`
}

// 签到记录
type SignIn struct {
	Id       int64     `json:"id" xorm:"pk"`
	UserId   int64     `json:"user_id" xorm:"index"`  //用户ID
	Username string    `json:"username" xorm:"index"` //账号
	Created  time.Time `json:"created" xorm:"created"`
}

// 聊天文本消息记录
type TextRecord struct {
	ChatText
	Type string `json:"type"`
	Time string `json:"time"`
}

// 聊天图片消息记录
type ImgRecord struct {
	Id      int64  `json:"id"`
	Room    string `json:"room"`
	ImgPath string `json:"img_path"`
	Type    string `json:"type"`
	Time    string `json:"time"`
}

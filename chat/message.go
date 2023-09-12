package chat

import "time"

// 加入
type Join struct {
	Id      int64  `json:"id"`
	Room    string `json:"room"`
	Content string `json:"content"`
}

// 文本
type Text struct {
	Id      int64  `json:"id"`
	Room    string `json:"room"`
	Content string `json:"content"`
}

// 图片
type Image struct {
	Id   int64  `json:"id"`
	Room string `json:"room"`
	//Img     []byte `json:"img"`
	ImgPath string `json:"img_path"`
}

// 红包记录
type RedPacket struct {
	Id           int64     `json:"id" xorm:"pk"`
	UserId       int64     `json:"user_id" xorm:"index"`
	Type         int64     `json:"type"` //0：默认红包，1:随机红包，
	Room         string    `json:"room"`
	CurrentMoney float64   `json:"current_money"`
	CurrentNum   int64     `json:"current_num"`
	TotalMoney   float64   `json:"total_money"`
	TotalNum     int64     `json:"total_num"`
	Created      time.Time `json:"created" xorm:"created"`
}

// 抢红包
type GrabPacket struct {
	Id      int64     `json:"id" xorm:"index"`
	UserId  int64     `json:"user_id" xorm:"index"`
	Money   float64   `json:"money"`
	Room    string    `json:"room"`
	Created time.Time `json:"created" xorm:"created"`
}

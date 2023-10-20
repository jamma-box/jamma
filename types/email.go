package types

import "time"

type Email struct {
	Id      int64     `json:"id,omitempty" form:"id"`
	UserId  string    `json:"user_id" form:"user_id"`           //目的地
	Title   string    `json:"title" form:"title"`               //主题
	Content string    `json:"content,omitempty" form:"content"` //邮件内容
	Read    bool      `json:"read" form:"read"`                 //是否已读，true表示已读
	Created time.Time `json:"created" xorm:"created" form:"created"`
}

package types

import "time"

type Email struct {
	Id      int64     `json:"id" xorm:"pk"`
	UserId  int64     `json:"user_id" xorm:"index"` //目的地
	Title   string    `json:"title"`                //主题
	Content string    `json:"content,omitempty"`    //邮件内容
	Read    bool      `json:"read"`                 //是否已读，true表示已读
	Created time.Time `json:"created" xorm:"created"`
}

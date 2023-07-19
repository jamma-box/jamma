package types

import "time"

type Email struct {
	Id      int64     `json:"id" xorm:"pk"`
	TO      string    `json:"to" `               //目的地
	From    string    `json:"from"`              //从哪发送
	Topic   string    `json:"topic"`             //主题
	File    any       `json:"file,omitempty"`    //附加文件
	Context string    `json:"context,omitempty"` //邮件内容
	Created time.Time `json:"created" xorm:"created"`
}

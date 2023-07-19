package types

import "time"

type GameHall struct {
	Id       int64     `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`     //名称
	Desc     string    `json:"desc,omitempty"`     //说明
	Img      string    `json:"img,omitempty"`      //图片
	Disabled bool      `json:"disabled,omitempty"` //禁用
	Created  time.Time `json:"created" xorm:"created"`
}

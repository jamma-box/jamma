package types

import "time"

type Game struct {
	Id       string    `json:"id,omitempty"  xorm:"pk"`
	Icon     string    `json:"icon,omitempty"` //图标
	Name     string    `json:"name,omitempty"` //名称
	Desc     string    `json:"desc,omitempty"` //说明
	Type     string    `json:"type,omitempty"` //类型
	Disabled bool      `json:"disabled"`       //禁用
	Created  time.Time `json:"created" xorm:"created"`
}

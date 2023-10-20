package types

import "time"

type Game struct {
	Id       int64     `json:"id,omitempty" form:"id"`
	Icon     string    `json:"icon,omitempty" form:"icon"` //图标
	Name     string    `json:"name,omitempty" form:"name"` //名称
	Desc     string    `json:"desc,omitempty" form:"desc"` //说明
	Type     string    `json:"type,omitempty" form:"type"` //类型
	Disabled bool      `json:"disabled" form:"disabled"`   //禁用
	Created  time.Time `json:"created" xorm:"created" form:"created"`
}

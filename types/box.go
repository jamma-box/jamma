package types

import "time"

type Box struct {
	Id       int64     `json:"id,omitempty" xorm:"pk"`
	Name     string    `json:"name,omitempty"`       //名称
	Desc     string    `json:"desc,omitempty"`       //说明
	Icon     string    `json:"icon,omitempty"`       //图片
	Type     string    `json:"type,omitempty"`       //类型
	Disabled bool      `json:"disabled"`             //禁用
	GameId   int64     `json:"game_id" xorm:"index"` //游戏机ID
	Created  time.Time `json:"created" xorm:"created"`
}

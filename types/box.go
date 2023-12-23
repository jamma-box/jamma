package types

import "time"

type Box struct {
	Id       string    `json:"id,omitempty" xorm:"pk" form:"id"`
	GameId   string    `json:"game_id" form:"game_id"`     //游戏机ID
	Name     string    `json:"name,omitempty" form:"name"` //名称
	Desc     string    `json:"desc,omitempty" form:"desc"` //说明
	Disabled bool      `json:"disabled" form:"disabled"`   //禁用
	Created  time.Time `json:"created" xorm:"created" form:"created"`
}

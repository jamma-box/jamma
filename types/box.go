package types

import "time"

type Box struct {
	Id       string    `json:"id,omitempty" xorm:"pk"`
	GameId   string    `json:"game_id" xorm:"index"` //游戏机ID
	Name     string    `json:"name,omitempty"`       //名称
	Desc     string    `json:"desc,omitempty"`       //说明
	Disabled bool      `json:"disabled"`             //禁用
	Created  time.Time `json:"created" xorm:"created"`
}

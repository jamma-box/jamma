package types

import "time"

type GameBox struct {
	Id      int64     `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Type    string    `json:"type,omitempty"`
	Created time.Time `json:"created" xorm:"created"`
}

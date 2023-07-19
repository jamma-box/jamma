package types

import "time"

type Image struct {
	Id      int64     `json:"id" xorm:"pk"`
	Name    string    `json:"name" xorm:"unique"`
	Data    []uint64  `json:"data"`
	Created time.Time `json:"created" xorm:"created"`
}

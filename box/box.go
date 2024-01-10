package box

import (
	"arcade/types"
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

type Seat struct {
	UserId int64           `json:"user_id"`
	LastId int64           `json:"last_id"`
	Client *websocket.Conn `json:"-"`
}

type Box struct {
	*types.Box
	Seats []Seat

	gamePad  *websocket.Conn
	gameLive *websocket.Conn

	//客户端接入
	client *websocket.Conn
	mt     int
}

func New(b *types.Box) *Box {
	return &Box{
		Box:   b,
		Seats: make([]Seat, 8),
	}
}

type PadCommand struct {
	Seat int    `json:"seat,omitempty"`
	Type string `json:"type,omitempty"`
	Coin int    `json:"coin,omitempty"`
}

func (b *Box) Bridge(c *websocket.Conn) {
	b.client = c

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Println("bridge", string(message))

		b.mt = mt
		if b.gameLive != nil {
			err = b.gameLive.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
	b.client = nil
}

func (b *Box) Live(c *websocket.Conn) {
	b.gameLive = c
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("live", string(message))

		if b.client != nil {
			err = b.client.WriteMessage(mt, message)
			if err != nil {
				b.client = nil
				log.Println(err)
			}
		}
	}
	b.gameLive = nil
}

func (b *Box) Pad(c *websocket.Conn) {
	b.gamePad = c
	//msg := make([]byte, 0)
	var msg []byte
	last := time.Now().UnixMilli()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("pad", string(message))

		//拼接历史数据
		now := time.Now().UnixMilli()
		if len(msg) > 0 && now-last < 500 {
			msg = append(msg, message...)
		} else {
			msg = message
		}
		last = now

		var cmd PadCommand
		err = json.Unmarshal(msg, &cmd)
		if err != nil {
			log.Println(err)
			continue
		}
		msg = nil //解析成功，清空

		//处理退分
		if cmd.Type == "refund" {
			log.Println("refund", cmd.Coin)
			var user types.User
			has, err := db.Engine.ID(b.Seats[cmd.Seat].LastId).Get(&user)
			if err != nil {
				log.Println(err)
				continue
			}
			if has {
				user.Balance = user.Balance + float64(cmd.Coin)
				_, _ = db.Engine.ID(user.Id).Cols("balance").Update(&user)
			}
		}
	}
	b.gamePad = nil
}

func (b *Box) Seat(seat int, c *websocket.Conn, user int64) {
	b.Seats[seat].UserId = user
	b.Seats[seat].LastId = user
	b.Seats[seat].Client = c

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Println("seat", string(message))

		b.mt = mt
		if b.gamePad != nil {
			err = b.gamePad.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}

	b.Seats[seat].Client = nil

	//意外退出
	if b.Seats[seat].UserId != 0 {
		//超时退出 1 分钟退币
		time.AfterFunc(time.Minute, func() {
			if b.Seats[seat].Client != nil {
				return
			}

			if b.gamePad != nil {
				_ = b.gamePad.WriteJSON(map[string]any{"seat": seat, "type": "click", "key": "refund"})
			}

			//超时退出 2 分钟站起来
			time.AfterFunc(time.Minute, func() {
				if b.Seats[seat].Client != nil {
					return
				}

				if b.gamePad != nil {
					_ = b.gamePad.WriteJSON(map[string]any{"seat": seat, "type": "stand"})
				}

				b.Seats[seat].UserId = 0
			})
		})

	}

}

var boxes lib.Map[Box]

func Get(id string) *Box {
	b := boxes.Load(id)
	if b == nil {
		var bb types.Box
		has, err := db.Engine.ID(id).Get(&bb)
		if err != nil {
			log.Println(err)
			return nil
		}
		//自动创建
		if !has {
			_, err := db.Engine.InsertOne(&bb)
			if err != nil {
				log.Println(err)
				return nil
			}
		}
		b = New(&bb)
	}
	boxes.Store(id, b)

	return b
}

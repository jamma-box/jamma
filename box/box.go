package box

import (
	"arcade/types"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"log"
)

type Seat struct {
	UserId int64
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
	Seat   int
	Type   string
	Refund int
}

func (b *Box) Bridge(c *websocket.Conn) {
	b.client = c

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		err = b.gameLive.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			return
		}
		b.mt = mt
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
		if b.client != nil {
			err = b.client.WriteMessage(mt, message)
			if err != nil {
				b.client = nil
				log.Println(err)
			}
		}
	}
}

func (b *Box) Pad(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		var cmd PadCommand
		err = json.Unmarshal(message, &cmd)
		if err != nil {
			log.Println(err)
			continue
		}

		//处理退分
		if cmd.Type == "refund" {
			var user types.User
			has, err := db.Engine.ID(b.Seats[cmd.Seat].UserId).Get(&user)
			if err != nil {
				log.Println(err)
				continue
			}
			if has {
				user.Balance = user.Balance + float64(cmd.Refund)
				_, _ = db.Engine.ID(user.Id).Cols("balance").Update(&user)
			}
		}
	}
}

func (b *Box) Seat(c *websocket.Conn, seat int) {

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		err = b.gamePad.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			return
		}
		b.mt = mt
	}
}

var boxes lib.Map[Box]

func Get(id string) *Box {
	return boxes.Load(id)
}

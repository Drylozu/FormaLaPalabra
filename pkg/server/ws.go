package server

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Game struct {
	Word    string
	Time    time.Duration
	Paused  bool
	Clients map[net.Addr]*websocket.Conn
	Timer   *time.Ticker
}

type Message struct {
	Word        string `json:"word"`
	StartTimer  bool   `json:"startTimer"`
	ExtraTime   int    `json:"extraTime"`
	SetTime     int    `json:"setTime"`
	TogglePause bool   `json:"pause"`
	Wrong       bool   `json:"wrong"`
	Correct     bool   `json:"correct"`
}

const gameTick = time.Second
const defaultDuration = 60 * gameTick

var screenWrong = map[string]string{
	"screen": "wrong",
}
var screenCorrect = map[string]string{
	"screen": "correct",
}

func NewGame() *Game {
	game := &Game{
		Word:    "",
		Time:    -1,
		Paused:  true,
		Clients: make(map[net.Addr]*websocket.Conn),
		Timer:   time.NewTicker(gameTick),
	}

	go func() {
		var (
			out []byte
			err error
		)

		for {
			if !game.Paused {
				game.Time -= gameTick
				if game.Time < gameTick {
					game.Paused = true
				}

				if out, err = sonic.Marshal(map[string]int{
					"time": int(game.Time / gameTick),
				}); err == nil {
					for _, conn := range game.Clients {
						conn.WriteMessage(1, out)
					}
				}
			}
			<-game.Timer.C
		}
	}()

	return game
}

func HandleWS(game *Game) func(*fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		game.Clients[c.RemoteAddr()] = c

		defer func() {
			delete(game.Clients, c.RemoteAddr())
		}()

		var (
			mt  int
			msg []byte
			err error

			out    []byte
			screen map[string]string
		)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				if websocket.IsUnexpectedCloseError(err) {
					c.Close()
					break
				}

				log.Println("err reading:", err)
				continue
			}

			var decoded Message
			if err = sonic.Unmarshal(msg, &decoded); err != nil {
				log.Println("err parsing:", err)
				continue
			}

			fmt.Printf("%#v\n", decoded)
			if decoded.StartTimer {
				game.Time = defaultDuration
				game.Paused = false
			}

			if decoded.SetTime != 0 {
				game.Time = time.Duration(decoded.SetTime) * gameTick
			}

			if decoded.ExtraTime != 0 {
				game.Time += time.Duration(decoded.ExtraTime) * gameTick
			}

			if decoded.Word != "" {
				game.Word = decoded.Word
			}

			if decoded.Wrong || decoded.Correct {
				if decoded.Wrong {
					screen = screenWrong
				} else {
					screen = screenCorrect
				}

				if out, err = sonic.Marshal(screen); err == nil {
					for _, conn := range game.Clients {
						conn.WriteMessage(mt, out)
					}
				}
			}

			if decoded.TogglePause {
				game.Paused = !game.Paused
			}
		}
	})
}

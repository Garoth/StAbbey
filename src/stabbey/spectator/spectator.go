package spectator

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type Spectator struct {
	WebSocketConnection *websocket.Conn
}

func New() *Spectator {
	return &Spectator{nil}
}

func (me *Spectator) GetWebSocketConnection() *websocket.Conn {
	return me.WebSocketConnection
}

func (me *Spectator) SetWebSocketConnection(conn *websocket.Conn) {
	me.WebSocketConnection = conn
}

func (me *Spectator) SendMessage(json string) error {
	if me.WebSocketConnection != nil {
		websocket.Message.Send(me.WebSocketConnection, json)
		return nil
	} else {
		return fmt.Errorf("Attempted to use websocket before ready")
	}
}

package interfaces

import (
	"golang.org/x/net/websocket"
)

type Spectator interface {
	/* Websocket Connection Getters / Setters */
	GetWebSocketConnection() *websocket.Conn
	SetWebSocketConnection(conn *websocket.Conn)
	/* Send a (json) message to the spectator */
	SendMessage(json string) error
}

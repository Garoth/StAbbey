package interfaces

import (
    "time"

    "code.google.com/p/go.net/websocket"
)

type Player interface {
    /* Unique Game ID Getters / Setters */
    GetPlayerId() int
    SetPlayerId(id int)
    /* Last Sent Tick Getters / Setters */
    GetLastTick() int
    SetLastTick(tickNum int)
    /* Last Sent Tick Time Getters / Setters */
    GetLastTickTime() time.Time
    SetLastTickTime(t time.Time)
    /* Queue Manipulation Getters / Setters */
    GetActionQueue() []Action
    GetStringActionQueue() []string
    SetActionQueue([]Action)
    PopAction() Action
    /* Websocket Connection Getters / Setters */
    GetWebSocketConnection() *websocket.Conn
    SetWebSocketConnection(conn *websocket.Conn)
    /* Send a (json) message to the player */
    SendMessage(json string)
}

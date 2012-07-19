package player

import (
    "time"

    "code.google.com/p/go.net/websocket"

    "stabbey/uidgenerator"
)

var uidg = uidgenerator.New();

type Player struct {
    PlayerId int
    PlayerLastTick int
    PlayerLastTickTime time.Time
    WebSocketConnection *websocket.Conn
}

func New() *Player {
    return &Player{uidg.NextUid(), 0, time.Now(), nil}
}

func (p *Player) GetPlayerId() int {
    return p.PlayerId
}

func (p *Player) SetPlayerId(id int) {
    p.PlayerId = id
}

func (p *Player) GetLastTick() int {
    return p.PlayerLastTick
}

func (p *Player) SetLastTick(tickNum int) {
    p.PlayerLastTick = tickNum
}

func (p *Player) GetLastTickTime() time.Time {
    return p.PlayerLastTickTime
}

func (p *Player) SetLastTickTime(t time.Time) {
    p.PlayerLastTickTime = t
}

func (p *Player) GetWebSocketConnection() *websocket.Conn {
    return p.WebSocketConnection
}

func (p *Player) SetWebSocketConnection(conn *websocket.Conn) {
    p.WebSocketConnection = conn
}

func (p *Player) SendMessage(json string) {
    websocket.Message.Send(p.WebSocketConnection, json)
}

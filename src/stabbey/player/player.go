package player

import (
    "strconv"
    "time"

    "code.google.com/p/go.net/websocket"

    "stabbey/constants"
    "stabbey/entity"
    "stabbey/uidgenerator"
)

var uidg = uidgenerator.New();

type Player struct {
    entity.Entity
    PlayerId int
    PlayerLastTick int
    PlayerLastTickTime time.Time
    WebSocketConnection *websocket.Conn
}

func New() *Player {
    p := &Player{}

    /* Player stuff */
    p.SetPlayerId(uidg.NextUid())
    p.SetLastTickTime(time.Now())
    p.SetWebSocketConnection(nil)

    /* Entity stuff */
    p.SetPosition(0, 8, 6)
    p.SetEntityId(entity.UIDG.NextUid())
    p.SetType(constants.ENTITY_TYPE_PLAYER)
    p.SetName("Player" + strconv.Itoa(p.GetPlayerId()))

    return p
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

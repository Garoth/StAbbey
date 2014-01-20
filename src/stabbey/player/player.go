package player

import (
    "log"
    "strconv"
    "time"

    "code.google.com/p/go.net/websocket"

    "stabbey/entity"
    "stabbey/interfaces"
    "stabbey/order"
    "stabbey/uidgenerator"
)

var uidg = uidgenerator.New()

type Player struct {
    *entity.Entity
    PlayerId int
    AvailableActions []interfaces.Action
    PlayerLastTick int
    PlayerLastTickTime time.Time
    WebSocketConnection *websocket.Conn
}

func New(g interfaces.Game) *Player {
    p := &Player{}

    /* Player stuff */
    p.SetPlayerId(uidg.NextUid())
    p.SetLastTickTime(time.Now())
    p.SetWebSocketConnection(nil)
    p.AvailableActions = append(p.AvailableActions, order.NewAction("."))
    p.AvailableActions = append(p.AvailableActions, order.NewAction("mu"))
    p.AvailableActions = append(p.AvailableActions, order.NewAction("*u"))
    p.AvailableActions = append(p.AvailableActions, order.NewAction("lu"))

    /* Entity stuff */
    p.Entity = entity.New(entity.UIDG.NextUid(), g)
    p.SetType(interfaces.ENTITY_TYPE_PLAYER)
    p.SetName("Player " + strconv.Itoa(p.GetPlayerId()))
    p.SetMaxArdour(100)
    p.SetArdour(100)

    return p
}

func (p *Player) GetPlayerId() int {
    return p.PlayerId
}

func (p *Player) SetPlayerId(id int) {
    p.PlayerId = id
}

func (p *Player) GetAvailableActions() []interfaces.Action {
    return p.AvailableActions
}

func (p *Player) AddAvailableAction(action interfaces.Action) {
    p.AvailableActions = append(p.AvailableActions, action)
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
    if p.WebSocketConnection != nil {
        websocket.Message.Send(p.WebSocketConnection, json)
    } else {
        log.Printf("Attempt to use player %v's websocket before its ready",
            p.GetPlayerId())
    }
}

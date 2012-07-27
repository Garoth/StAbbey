package player

import (
    "strconv"
    "time"

    "code.google.com/p/go.net/websocket"

    "stabbey/constants"
    "stabbey/entity"
    "stabbey/uidgenerator"
    "stabbey/interfaces"
)

var uidg = uidgenerator.New();

type Player struct {
    entity.Entity
    PlayerId int
    PlayerLastTick int
    PlayerLastTickTime time.Time
    ActionQueue []interfaces.Action
    WebSocketConnection *websocket.Conn
}

func New() *Player {
    p := &Player{}

    /* Player stuff */
    p.SetPlayerId(uidg.NextUid())
    p.SetLastTickTime(time.Now())
    p.SetWebSocketConnection(nil)
    p.ActionQueue = make([]interfaces.Action, 0, 10)

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

func (p *Player) GetActionQueue() []interfaces.Action {
    return p.ActionQueue
}

func (p *Player) GetStringActionQueue() []string {
    q := make([]string, len(p.GetActionQueue()))

    for i := 0; i < len(p.GetActionQueue()); i++ {
        q[i] = p.GetActionQueue()[i].ActionType()
    }

    return q
}

func (p *Player) SetActionQueue(aq []interfaces.Action) {
    p.ActionQueue = aq
}

func (p *Player) PopAction() interfaces.Action {
    if len(p.ActionQueue) > 0 {
        a := p.ActionQueue[0]
        p.ActionQueue = p.ActionQueue[1:]
        return a
    }
    return nil
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

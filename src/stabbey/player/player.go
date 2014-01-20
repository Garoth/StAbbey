package player

import (
    "strconv"
    "time"

    "stabbey/spectator"
    "stabbey/entity"
    "stabbey/interfaces"
    "stabbey/order"
    "stabbey/uidgenerator"
)

var uidg = uidgenerator.New()

type Player struct {
    *spectator.Spectator
    *entity.Entity
    PlayerId int
    AvailableActions []interfaces.Action
    PlayerLastTick int
    PlayerLastTickTime time.Time
}

func New(g interfaces.Game) *Player {
    p := &Player{}

    /* Spectator stuff */
    p.Spectator = spectator.New()

    /* Player stuff */
    p.SetPlayerId(uidg.NextUid())
    p.SetLastTickTime(time.Now())
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

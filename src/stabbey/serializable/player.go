package serializable

import (
    "stabbey/interfaces"
)

type Player struct {
    /* Player Data */
    Id int
    EntityId int
    AvailableActions []*Action
}

func NewPlayer(p interfaces.Player) *Player {
    me := &Player{}

    me.Id = p.GetPlayerId()
    me.EntityId = p.GetEntityId()

    actions := p.GetAvailableActions()
    for k := 0; k < len(actions); k++ {
        me.AvailableActions = append(me.AvailableActions, NewAction(actions[k]))
    }

    return me
}

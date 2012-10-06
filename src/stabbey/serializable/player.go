package serializable

import (
    "stabbey/interfaces"
)

type Player struct {
    /* Player Data */
    Id int
    EntityId int
}

func NewPlayer(p interfaces.Player) *Player {
    return &Player{p.GetPlayerId(), p.GetEntityId()}
}

package serializable

import (
    "stabbey/interfaces"
)

type Player struct {
    /* Player Data */
    Id int
    ActionQueue []string
}

func NewPlayer(p interfaces.Player) *Player {
    return &Player{p.GetPlayerId(), p.GetStringActionQueue()}
}

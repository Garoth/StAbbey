package stabbey

import (
)

type SerializablePlayer struct {
    /* Player Data */
    Id string
    /* Entity Data */
    EntityId, X, Y int
    Name string
}

func NewSerializablePlayer(p *Player) *SerializablePlayer {
    sp := &SerializablePlayer{}

    sp.Id = p.Id
    sp.EntityId = p.EntityId
    sp.X = p.X
    sp.Y = p.Y
    sp.Name = p.Name

    return sp
}

package stabbey

import (
)

type SerializablePlayer struct {
    /* Player Data */
    Id int
    /* Entity Data */
    EntityId int
    Name string
}

func NewSerializablePlayer(p *Player) *SerializablePlayer {
    sp := &SerializablePlayer{}

    sp.Id = p.Id
    sp.EntityId = p.EntityId
    sp.Name = p.Name

    return sp
}

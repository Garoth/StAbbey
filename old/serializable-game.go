package stabbey

import (
)

/* Object used for JSON serialization */
type SerializableGame struct {
    Players []*SerializablePlayer
    Boards []*SerializableBoard
    Entities []*SerializableEntity
    LastTick int
}

func NewSerializableGame(c *Context, game *Game) *SerializableGame {
    sg := &SerializableGame{}

    for _, ID := range game.Players {
        sg.Players = append(sg.Players, NewSerializablePlayer(LoadPlayer(c, ID)))
    }

    for _, ID := range game.Boards {
        sg.Boards = append(sg.Boards, NewSerializableBoard(LoadBoard(c, ID)))
    }

    for _, ID := range game.Entities {
        sg.Entities = append(sg.Entities, NewSerializableEntity(LoadEntity(c, ID)))
    }

    sg.LastTick = game.LastTick

    return sg
}

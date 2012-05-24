package stabbey

import (
)

/* Object used for JSON serialization */
type SerializableGame struct {
    Players []*SerializablePlayer
    Boards []*Board
    LastTick int
}

func NewSerializableGame(c *Context, game *Game) *SerializableGame {
    sg := &SerializableGame{}

    for _, ID := range game.Players {
        sg.Players = append(sg.Players, NewSerializablePlayer(LoadPlayer(c, ID)))
    }

    for _, ID := range game.Boards {
        sg.Boards = append(sg.Boards, LoadBoard(c, string(ID)))
    }

    sg.LastTick = game.LastTick

    return sg
}

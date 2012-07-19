package serializable

import (
    "stabbey/interfaces"
)

/* Object used for JSON serialization */
type Game struct {
    Players []*Player
    Boards []*Board
    Entities []*Entity
    LastTick int
}

func NewGame(game interfaces.Game) *Game {
    g := &Game{}

    for _, player := range game.GetPlayers() {
        g.Players = append(g.Players, NewPlayer(player))
    }
    for _, board := range game.GetBoards() {
        g.Boards = append(g.Boards, NewBoard(board))
    }
    for _, entity := range game.GetEntities() {
        g.Entities = append(g.Entities, NewEntity(entity))
    }

    g.LastTick = game.GetLastTick()

    return g
}

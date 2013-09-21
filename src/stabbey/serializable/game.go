package serializable

import (
    "stabbey/interfaces"
)

/* Object used for JSON serialization */
type Game struct {
    Players []*Player
    Boards []*Board
    Entities []*Entity
    LastTick, CurrentBoard int
}

func NewGame(game interfaces.Game) *Game {
    g := &Game{}

    for k := 0; k < len(game.GetPlayers()); k++ {
        g.Players = append(g.Players, NewPlayer(game.GetPlayers()[k]))
    }
    for k := 0; k < len(game.GetBoards()); k++ {
        g.Boards = append(g.Boards, NewBoard(game.GetBoards()[k]))
    }
    for k := 0; k < len(game.GetEntities()); k++ {
        g.Entities = append(g.Entities, NewEntity(game.GetEntities()[k]))
    }

    g.LastTick = game.GetLastTick()
    g.CurrentBoard = game.GetCurrentBoard()

    return g
}

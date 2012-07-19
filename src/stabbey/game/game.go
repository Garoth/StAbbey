package game

import (
    "log"
    "stabbey/interfaces"
    "stabbey/util"
)

type Game struct {
    Players map[int] interfaces.Player
    Boards map[int] interfaces.Board
    Entities map[int] interfaces.Entity
    LastTick int
    GameRunning bool
    Gamekey string
}

func NewGame(gamekey string) *Game {
    log.Printf("Starting new game, %v", gamekey)
    g := &Game{}
    g.GameRunning = false
    g.Gamekey = gamekey
    g.Players = make(map[int] interfaces.Player, 10)
    g.Boards = make(map[int] interfaces.Board, 10)
    g.Entities = make(map[int] interfaces.Entity, 100)
    return g
}

func (g *Game) AddPlayer(player interfaces.Player) {
    g.Players[player.GetPlayerId()] = player
}

func (g *Game) GetPlayer(id int) interfaces.Player {
    return g.Players[id]
}

func (g *Game) AddBoard(board interfaces.Board) {
    g.Boards[board.GetLevel()] = board
}

func (g *Game) GetBoard(level int) interfaces.Board {
    return g.Boards[level]
}

func (g *Game) AddEntity(entity interfaces.Entity) {
    g.Entities[entity.GetEntityId()] = entity
}

func (g *Game) GetEntity(entid int) interfaces.Entity {
    return g.Entities[entid]
}

/* Generates the gamestate for the given player's perspective */
func (g *Game) Json(playerId int) string {
    util.Stub("g.Json")
    return ""
}

func (g *Game) Run() {
    log.Printf("Game %v running", g.Gamekey)
}

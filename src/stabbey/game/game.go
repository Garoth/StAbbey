package game

import (
    "encoding/json"
    "log"

    "stabbey/interfaces"
    "stabbey/serializable"
)

type Game struct {
    Players map[int] interfaces.Player
    Boards map[int] interfaces.Board
    Entities map[int] interfaces.Entity
    PlayersToEntities map[interfaces.Player] interfaces.Entity
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
    g.PlayersToEntities = make(map[interfaces.Player] interfaces.Entity, 10)
    return g
}

func (g *Game) AddPlayer(player interfaces.Player, entity interfaces.Entity) {
    g.Players[player.GetPlayerId()] = player
    g.AddEntity(entity)
    g.PlayersToEntities[player] = entity
}

func (g *Game) GetPlayer(id int) interfaces.Player {
    return g.Players[id]
}

func (g *Game) GetPlayers() map[int] interfaces.Player {
    return g.Players
}

func (g *Game) AddBoard(board interfaces.Board) {
    g.Boards[board.GetLevel()] = board
}

func (g *Game) GetBoard(level int) interfaces.Board {
    return g.Boards[level]
}

func (g *Game) GetBoards() map[int] interfaces.Board {
    return g.Boards
}

func (g *Game) AddEntity(entity interfaces.Entity) {
    g.Entities[entity.GetEntityId()] = entity
}

func (g *Game) GetEntity(entid int) interfaces.Entity {
    return g.Entities[entid]
}

func (g *Game) GetEntityByPlayer(player interfaces.Player) interfaces.Entity {
    return g.PlayersToEntities[player]
}

func (g *Game) GetEntities() map[int] interfaces.Entity {
    return g.Entities
}

func (g *Game) GetLastTick() int {
    return g.LastTick
}

func (g *Game) SetLastTick(tick int) {
    g.LastTick = tick
}

/* Generates the gamestate for the given player's perspective */
func (g *Game) Json(player interfaces.Player) string {
    b, e := json.Marshal(serializable.NewGame(g))
    if e != nil {
        log.Fatalf("Error serializing game: %v", e)
    }
    return string(b)
}

func (g *Game) Run() {
    log.Printf("Game %v running", g.Gamekey)
}

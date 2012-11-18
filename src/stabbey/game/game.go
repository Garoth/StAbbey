package game

import (
    "encoding/json"
    "log"
    "math/rand"
    "time"

    "stabbey/interfaces"
    "stabbey/serializable"
)

type Game struct {
    Players map[int] interfaces.Player
    Boards map[int] interfaces.Board
    Entities map[int] interfaces.Entity
    PlayersToEntities map[interfaces.Player] interfaces.Entity
    EntityIdToMonster map[int] interfaces.Monster
    LastTick int
    CurrentBoard int
    GameRunning bool
    Gamekey string
}

func NewGame(gamekey string) *Game {
    log.Printf("Starting new game, %v", gamekey)
    rand.Seed(time.Now().Unix())
    g := &Game{}
    g.GameRunning = false
    g.Gamekey = gamekey
    g.Players = make(map[int] interfaces.Player, 10)
    g.Boards = make(map[int] interfaces.Board, 10)
    g.CurrentBoard = 0
    g.Entities = make(map[int] interfaces.Entity, 100)
    g.PlayersToEntities = make(map[interfaces.Player] interfaces.Entity, 10)
    g.EntityIdToMonster = make(map[int] interfaces.Monster, 10)
    return g
}

func (g *Game) AddPlayer(player interfaces.Player, entity interfaces.Entity) {
    g.Players[player.GetPlayerId()] = player
    g.AddEntity(entity)
    g.PlayersToEntities[player] = entity
}

func (g *Game) AddMonster(mon interfaces.Monster) {
    g.EntityIdToMonster[mon.GetEntityId()] = mon
    g.AddEntity(mon)
}

/* Checks if an entity may be placed into a given tile on the current board */
func (g *Game) IsSpaceEmpty(locX, locY int) bool {
    /* Check if any entities are there */
    for _, entity := range g.Entities {
        boardId, x, y := entity.GetPosition()
        if boardId == g.CurrentBoard && x == locX && y == locY {
            return false
        }
    }

    /* Check if there are walls there */
    /* TODO this is both inefficient and fragile */
    if (g.Boards[g.CurrentBoard].GetRender()[locY][locX] == '#') {
        return false
    }

    return true
}

/* Picks a random empty space */
func (g *Game) GetRandomEmptySpace() (int, int) {
    maxAttempts := 1000

    for x := 0; x < maxAttempts; x++ {
        x := rand.Intn(interfaces.BOARD_WIDTH)
        y := rand.Intn(interfaces.BOARD_HEIGHT)

        if g.IsSpaceEmpty(x, y) {
            return x, y
        }
    }

    /* TODO should never happen */
    return 0, 0
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

func (g *Game) GetMonsterByEntityId(entid int) interfaces.Monster {
    return g.EntityIdToMonster[entid]
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

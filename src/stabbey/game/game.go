package game

/* Represents the overall game state, and hooks together all the various parts:
 *  - Boards
 *  - Entities
 *  - Players
 *  - Monsters
 *  - Etc
 *
 * Can be queried somewhat like a database to find out about the game / change
 * its state.
 */

/* FIXME This likely needs some big scary global-lock style mutexes, since
 *       this is a central object that's referred to all over the place. I
 *       can imagine nasty race conditions happening in the long run, so it's
 *       probably best to just lock it down frequently.
 */

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
    return g
}

func (g *Game) AddPlayer(player interfaces.Player, entity interfaces.Entity) {
    g.Players[player.GetPlayerId()] = player
    g.AddEntity(entity)
    g.PlayersToEntities[player] = entity
}

/* Checks if an entity may be placed into a given tile on the current board */
func (g *Game) CanMoveToSpace(locX, locY int) bool {
    /* Check if any entities are there */
    for _, e := range g.Entities {
        boardId, x, y := e.GetPosition()
        if boardId == g.CurrentBoard && x == locX && y == locY &&
                e.IsTangible() {
            return false
        }
    }

    /* Check if there are walls there */
    if g.IsWall(locX, locY) {
        return false
    }

    return true
}

/* Checks if there's a wall at a given tile */
func (g *Game) IsWall(x, y int) bool {
    if x > interfaces.BOARD_WIDTH - 1|| y > interfaces.BOARD_HEIGHT - 1 ||
            x < 0 || y < 0 {
        log.Println("Tried to check out wall out of bounds")
        return true
    }

    /* TODO this is both inefficient and fragile */
    if (g.Boards[g.CurrentBoard].GetRender()[y][x] == '#') {
        return true
    }

    return false
}

/* Checks if there's water at a given tile */
func (g *Game) IsWater(x, y int) bool {
    if x > interfaces.BOARD_WIDTH - 1|| y > interfaces.BOARD_HEIGHT - 1 ||
            x < 0 || y < 0 {
        log.Println("Tried to check for water out of bounds")
        return false
    }

    /* TODO this is both inefficient and fragile */
    if (g.Boards[g.CurrentBoard].GetRender()[y][x] == '~') {
        return true
    }

    return false
}

/* Picks a random empty space */
func (g *Game) GetRandomEmptySpace() (int, int) {
    maxAttempts := 1000

    for x := 0; x < maxAttempts; x++ {
        x := rand.Intn(interfaces.BOARD_WIDTH)
        y := rand.Intn(interfaces.BOARD_HEIGHT)

        /* Has to be somewhere an entity can stand, and not in water */
        if g.CanMoveToSpace(x, y) && !g.IsWater(x, y) {
            return x, y
        }
    }

    /* TODO should never happen */
    return 0, 0
}


func (g *Game) GetPlayer(id int) interfaces.Player {
    return g.Players[id]
}

func (g *Game) GetPlayerByEntity(entity interfaces.Entity) interfaces.Player {
    for player, ent := range g.PlayersToEntities {
        if ent.GetEntityId() == entity.GetEntityId() {
            return player
        }
    }

    return nil
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

func (g *Game) GetEntitiesAtSpace(boardId, x, y int) []interfaces.Entity {
    ret := []interfaces.Entity{}

    for _, entity := range g.Entities {
        boardId2, x2, y2 := entity.GetPosition()
        if boardId2 == boardId && x == x2 && y == y2 {
            ret = append(ret, entity)
        }
    }

    return ret
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

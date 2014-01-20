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
    "stabbey/board"
    "stabbey/serializable"
    "stabbey/entity"
)

type Game struct {
    Players map[int] interfaces.Player
    Boards map[int] interfaces.Board
    Entities map[int] interfaces.Entity
    PlayersToEntities map[interfaces.Player] interfaces.Entity
    LastTick, CurrentBoard, NumBoards int
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
    g.Entities = make(map[int] interfaces.Entity, 100)
    g.PlayersToEntities = make(map[interfaces.Player] interfaces.Entity, 10)
    g.Boards = make(map[int] interfaces.Board, 10)
    g.CurrentBoard = -1
    g.NumBoards = 3
    /* Do at the end, to ensure game object is ready for modification */
    for i := 0; i < g.NumBoards; i++ {
        g.AddBoard(board.New(i, g))
        g.GetBoards()[i].LoadStartingEntities()
    }
    g.NextBoard()
    return g
}

func (g *Game) AddPlayer(player interfaces.Player, entity interfaces.Entity) {
    g.Players[player.GetPlayerId()] = player
    g.AddEntity(entity)
    g.PlayersToEntities[player] = entity
}

/* Checks if an entity may be placed into a given tile on the current board */
func (g *Game) CanMoveToSpace(targetBoardId, locX, locY int) bool {

    /* Check possible board */
    if targetBoardId >= len(g.Boards) {
        return false
    }

    if locX < 0 || locX >= interfaces.BOARD_WIDTH ||
            locY < 0 || locY >= interfaces.BOARD_HEIGHT {
        return false
    }

    /* Check if any entities are there */
    for _, e := range g.Entities {
        boardId, x, y := e.GetPosition()
        if boardId == targetBoardId && x == locX && y == locY &&
                e.IsTangible() {
            return false
        }
    }

    /* Check if there are walls there */
    if g.IsWall(targetBoardId, locX, locY) {
        return false
    }

    return true
}

/* Checks if there's a wall at a given tile */
func (g *Game) IsWall(boardId, x, y int) bool {
    if x > interfaces.BOARD_WIDTH - 1|| y > interfaces.BOARD_HEIGHT - 1 ||
            x < 0 || y < 0 {
        return true
    }

    /* TODO this is both inefficient and fragile */
    if (g.Boards[boardId].GetRender()[y][x] == '#') {
        return true
    }

    return false
}

/* Checks if there's water at a given tile */
func (g *Game) IsWater(boardId, x, y int) bool {
    if x > interfaces.BOARD_WIDTH - 1|| y > interfaces.BOARD_HEIGHT - 1 ||
            x < 0 || y < 0 {
        log.Println("Tried to check for water out of bounds")
        return false
    }

    /* TODO this is both inefficient and fragile */
    if (g.Boards[boardId].GetRender()[y][x] == '~') {
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
        if g.CanMoveToSpace(g.GetCurrentBoard(), x, y) &&
                !g.IsWater(g.GetCurrentBoard(), x, y) {
            return x, y
        }
    }

    /* TODO should never happen */
    return 0, 0
}

func (g *Game) PlaceAtNearestTile(ent interfaces.Entity, boardId, x, y int) {
    tryX, tryY := -1, -1

    /* assuming BOARD_WIDTH is the higher of the two */
    for radius := 0; radius < interfaces.BOARD_WIDTH + 5; radius++ {

        /* Testing horizontally an vertically closest tiles */
        /* Right */
        tryX, tryY = x + radius, y
        if (g.CanMoveToSpace(boardId, tryX, tryY)) {
            ent.SetPosition(boardId, tryX, tryY)
            return
        }

        /* Bottom */
        tryX, tryY = x, y + radius
        if (g.CanMoveToSpace(boardId, tryX, tryY)) {
            ent.SetPosition(boardId, tryX, tryY)
            return
        }

        /* Left */
        tryX, tryY = x - radius, y
        if (g.CanMoveToSpace(boardId, tryX, tryY)) {
            ent.SetPosition(boardId, tryX, tryY)
            return
        }

        /* Top */
        tryX, tryY = x, y - radius
        if (g.CanMoveToSpace(boardId, tryX, tryY)) {
            ent.SetPosition(boardId, tryX, tryY)
            return
        }

        /* Testing all tiles at increasing distances away */
        for wallLen := 0; wallLen < (radius * 2) + 1; wallLen++ {
            /* Testing top wall of new circle */
            tryX, tryY = x - radius + wallLen, y - radius
            if (g.CanMoveToSpace(boardId, tryX, tryY)) {
                ent.SetPosition(boardId, tryX, tryY)
                return
            }

            /* Testing right wall of new circle */
            tryX, tryY = x + radius, y - radius + wallLen
            if (g.CanMoveToSpace(boardId, tryX, tryY)) {
                ent.SetPosition(boardId, tryX, tryY)
                return
            }

            /* Testing left wall of new circle */
            tryX, tryY = x - radius, y - radius + wallLen
            if (g.CanMoveToSpace(boardId, tryX, tryY)) {
                ent.SetPosition(boardId, tryX, tryY)
                return
            }

            /* Testing bottom wall of new circle */
            tryX, tryY = x - radius + wallLen, y + radius
            if (g.CanMoveToSpace(boardId, tryX, tryY)) {
                ent.SetPosition(boardId, tryX, tryY)
                return
            }
        }
    }
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
    g.Boards[board.GetId()] = board
}

func (g *Game) GetBoard(boardId int) interfaces.Board {
    return g.Boards[boardId]
}

func (g *Game) GetBoards() map[int] interfaces.Board {
    return g.Boards
}

func (g *Game) GetCurrentBoard() int {
    return g.CurrentBoard
}

func (g *Game) GetNumBoards() int {
    return g.NumBoards
}

func (g *Game) placeEntityRandomly(entity interfaces.Entity, boardId int) {
    x, y := g.GetRandomEmptySpace()
    entity.SetPosition(boardId, x, y)
    g.AddEntity(entity)
}

func (g *Game) NextBoard() {
    if (g.CurrentBoard + 1 >= g.NumBoards) {
        log.Fatalf("Call to move to next board despite being at last one")
    }

    g.CurrentBoard++
    g.Boards[g.CurrentBoard].WarpPlayersToStart()

    /* Place some traps */
    for i := 0; i < 3; i++ {
        x, y := g.GetRandomEmptySpace()
        g.placeEntityRandomly(entity.NewTeleportTrap(g, x, y), g.CurrentBoard)
    }

    for i := 0; i < 3; i++ {
        g.placeEntityRandomly(entity.NewCaltropTrap(g), g.CurrentBoard)
    }
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

func (g *Game) GetTangibleEntityAtSpace(bId, x, y int) interfaces.Entity {
    /* A check to ensure that there is only one tangible entity would be
     * useful, since that should *never* happen */
    for _, entity := range g.GetEntitiesAtSpace(bId, x, y) {
        if entity.IsTangible() {
            return entity
        }
    }

    return nil
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
    g.Boards[g.CurrentBoard].WarpPlayersToStart()
    log.Printf("Game %v running", g.Gamekey)
}

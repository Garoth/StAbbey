package interfaces

type Game interface {
    /*** Spectator Manipulation ***/
    AddSpectator(spectator Spectator)
    GetSpectators() map[int] Spectator

    /*** Player Manipulation ***/
    AddPlayer(player Player, entity Entity)
    GetPlayer(id int) Player
    GetPlayerByEntity(entity Entity) Player
    GetPlayers() map[int] Player

    /*** Board manipulation ***/
    AddBoard(board Board)
    GetBoard(level int) Board
    GetBoards() map[int] Board
    GetCurrentBoard() int
    // Returns the amount of boards the game has
    GetNumBoards() int
    // Advances game to next stage, warps all players
    NextBoard()

    /*** Entity manipulation ***/
    AddEntity(entity Entity)
    GetEntity(entid int) Entity
    GetEntityByPlayer(player Player) Entity
    GetEntities() map[int] Entity
    GetEntitiesAtSpace(boardid, x, y int) []Entity
    // Returns first tangible entity at space (there should only ever be one)
    GetTangibleEntityAtSpace(boardid, x, y int) Entity
    IsWall(boardId, x, y int) bool
    IsWater(boardId, x, y int) bool

    /*** Utilities ***/
    // Checks if an entity may move to a given space
    CanMoveToSpace(boardId, x, y int) bool
    // Gets a random empty space for an entity to stand on
    GetRandomEmptySpace() (int, int)
    // Searches around given tile to find nearest empty space
    PlaceAtNearestTile(e Entity, boardId, x, y int)

    /*** Tick manipulation ***/
    GetLastTick() int
    SetLastTick(tick int)

    /*** Get the JSON representation ***/
    Json(player Player) string

    /*** Run the game (only launched by main) ***/
    Run()
}

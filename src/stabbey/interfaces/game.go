package interfaces

type Game interface {
    /* Player Manipulation */
    AddPlayer(player Player, entity Entity)
    GetPlayer(id int) Player
    GetPlayerByEntity(entity Entity) Player
    GetPlayers() map[int] Player
    /* Board manipulation */
    AddBoard(board Board)
    GetBoard(level int) Board
    GetBoards() map[int] Board
    GetCurrentBoard() int
    /* Entity manipulation */
    AddEntity(entity Entity)
    GetEntity(entid int) Entity
    GetEntityByPlayer(player Player) Entity
    GetEntities() map[int] Entity
    GetEntitiesAtSpace(boardid, x, y int) []Entity
    IsWall(boardId, x, y int) bool
    IsWater(boardId, x, y int) bool
    /* Utilities */
    CanMoveToSpace(boardId, x, y int) bool
    GetRandomEmptySpace() (int, int)
    PlaceAtNearestTile(e Entity, boardId, x, y int)
    /* Tick manipulation */
    GetLastTick() int
    SetLastTick(tick int)
    /* Get the JSON representation */
    Json(player Player) string
    /* Run the game (only launched by main) */
    Run()
}

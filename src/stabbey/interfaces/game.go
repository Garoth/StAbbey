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
    /* Entity manipulation */
    AddEntity(entity Entity)
    GetEntity(entid int) Entity
    GetEntityByPlayer(player Player) Entity
    GetEntityByLocation(boardId, x, y int) Entity
    GetEntities() map[int] Entity
    GetEntitiesAtSpace(boardid, x, y int) []Entity
    /* Monster manipulation */
    AddMonster(monster Monster)
    GetMonsterByEntityId(entid int) Monster
    /* Utilities */
    CanMoveToSpace(x, y int) bool
    GetRandomEmptySpace() (int, int)
    /* Tick manipulation */
    GetLastTick() int
    SetLastTick(tick int)
    /* Get the JSON representation */
    Json(player Player) string
    /* Run the game (only launched by main) */
    Run()
}

package interfaces

type Game interface {
    /* Player Manipulation */
    AddPlayer(player Player)
    GetPlayer(id int) Player
    GetPlayers() map[int] Player
    /* Board manipulation */
    AddBoard(board Board)
    GetBoard(level int) Board
    GetBoards() map[int] Board
    /* Entity manipulation */
    AddEntity(entity Entity)
    GetEntity(entid int) Entity
    GetEntities() map[int] Entity
    /* Tick manipulation */
    GetLastTick() int
    SetLastTick(tick int)
    /* Get the JSON representation */
    Json(player Player) string
    /* Run the game (only launched by main) */
    Run()
}

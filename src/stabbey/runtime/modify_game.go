package runtime

import (
    "log"
    "time"

    "stabbey/interfaces"
)

/* Checks whether all players are ready (done queueing) */
func allPlayersReady() bool {
    for _, player := range GAME.GetPlayers() {
        /* Ready players are 1 ahead of the game's tick */
        if player.GetLastTick() <= GAME.GetLastTick() {
            return false
        }
    }
    return true
}

/* Sends the gamestate to everyone */
func broadcastGamestate() {
    for _, player := range GAME.GetPlayers() {
        player.SendMessage(GAME.Json(player))
    }
}

/* Generic action handler for any entity */
func act(entity interfaces.Entity, action interfaces.Action) {
    command := action.ActionType()

    boardId, x, y := entity.GetPosition()
    if command == "mr" {
        entity.SetPosition(boardId, x + 1, y)
    } else if command == "ml" {
        entity.SetPosition(boardId, x - 1, y)
    } else if command == "mu" {
        entity.SetPosition(boardId, x, y - 1)
    } else if command == "md" {
        entity.SetPosition(boardId, x, y + 1)
    } else {
        log.Printf("Unknown order %v ignored!", command)
    }
}

/* Updates players' ticks and send out gamestate when everyone's ready */
func updateTick(order interfaces.Order) {
    p := order.GetPlayer()
    p.SetLastTick(order.GetTickNumber())
    p.SetLastTickTime(time.Now())

}
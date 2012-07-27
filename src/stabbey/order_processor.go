package main

import (
    "log"
    "time"

    "stabbey/interfaces"
)

var ORDER_STREAM = make(chan interfaces.Order)
var COMMAND_CODES = map[int]string {
    0: "start game",
    1: "update tick",
    2: "set queue" }

func ProcessOrders() {
    for {
        order := <-ORDER_STREAM

        if COMMAND_CODES[order.GetCommandCode()] == "update tick" {
            UpdateTick(order)
        } else if COMMAND_CODES[order.GetCommandCode()] == "set queue" {
            order.GetPlayer().SetActionQueue(order.GetActions())
        }

        log.Printf("Parsed Order: command:'%v' tick:%v actions:%v player:%v",
            COMMAND_CODES[order.GetCommandCode()], order.GetTickNumber(),
            order.GetPlayer().GetStringActionQueue(),
            order.GetPlayer().GetPlayerId())
    }
}

/* Updates players' ticks and send out gamestate when everyone's ready */
func UpdateTick(order interfaces.Order) {
    p := order.GetPlayer()
    p.SetLastTick(order.GetTickNumber())
    p.SetLastTickTime(time.Now())

    for _, player := range GAME.GetPlayers() {
        /* Ready players are 1 ahead of the game's tick */
        if player.GetLastTick() <= GAME.GetLastTick() {
            /* Not everyone's ready */
            return
        }
    }

    for _, player := range GAME.GetPlayers() {
        /* TODO this is fragile code that relies on playerid == entityid */
        entity := GAME.GetEntity(player.GetPlayerId())

        if entity == nil {
            log.Fatal("Got nil entity for Player %v", player.GetPlayerId())
        }

        if action := player.PopAction(); action != nil {
            MoveEntity(entity, action)
        } else {
            log.Printf("Player %v has empty queue", player.GetPlayerId())
        }
    }

    log.Printf("All players are ready, sending next tick")
    GAME.SetLastTick(GAME.GetLastTick() + 1)
    BroadcastGamestate()
}

/* Moves the entity according to the given action */
/* TODO really want a more general version of this that handles any action */
func MoveEntity(entity interfaces.Entity, action interfaces.Action) {
    command := action.ActionType()

    bid, x, y := entity.GetPosition()
    if command == "mr" {
        entity.SetPosition(bid, x + 1, y)
    } else if command == "ml" {
        entity.SetPosition(bid, x - 1, y)
    } else if command == "mu" {
        entity.SetPosition(bid, x, y - 1)
    } else if command == "md" {
        entity.SetPosition(bid, x, y + 1)
    } else {
        log.Printf("Unknown order %v ignored!", command)
    }
}

/* Sends the gamestate to everyone */
func BroadcastGamestate() {
    for _, player := range GAME.GetPlayers() {
        player.SendMessage(GAME.Json(player))
    }
}

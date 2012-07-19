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
        log.Printf("Found Order: command:%v tick:%v actions:%v player:%v",
            order.GetCommandCode(), order.GetTickNumber(), order.GetActions(),
            order.GetPlayer().GetPlayerId())

        if COMMAND_CODES[order.GetCommandCode()] == "update tick" {
            UpdateTick(order)
        } else if COMMAND_CODES[order.GetCommandCode()] == "set queue" {
            SetQueue(order)
        }
    }
}

/* Updates players' ticks and send out gamestate when everyone's ready */
func UpdateTick(order interfaces.Order) {
    p := order.GetPlayer()
    p.SetLastTick(order.GetTickNumber())
    p.SetLastTickTime(time.Now())

    allReady := true
    for _, player := range GAME.GetPlayers() {
        /* Ready players are 1 ahead of the game's tick */
        if player.GetLastTick() <= GAME.GetLastTick() {
            allReady = false
            break
        }
    }

    if allReady {
        log.Printf("All players are ready, sending next tick")
        GAME.SetLastTick(GAME.GetLastTick() + 1)
        BroadcastGamestate()
    }
}

/* Updates the player's queue of actions */
/* TODO this function is just a demo */
func SetQueue(order interfaces.Order) {
    /* TODO this is fragile code that relies on playerid == entityid */
    entity := GAME.GetEntity(order.GetPlayer().GetPlayerId())
    command := order.GetActions()[0]

    bid, x, y := entity.GetPosition()
    if command == "mr" {
        entity.SetPosition(bid, x + 1, y)
    } else if command == "ml" {
        entity.SetPosition(bid, x - 1, y)
    } else if command == "mu" {
        entity.SetPosition(bid, x, y - 1)
    } else if command == "md" {
        entity.SetPosition(bid, x, y + 1)
    }

    BroadcastGamestate()
}

/* Sends the gamestate to everyone */
func BroadcastGamestate() {
    for _, player := range GAME.GetPlayers() {
        player.SendMessage(GAME.Json(player))
    }
}

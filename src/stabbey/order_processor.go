package main

import (
    "log"
    "time"

    "stabbey/interfaces"
    "stabbey/runtime"
)

var ORDER_STREAM = make(chan interfaces.Order)
var RUNTIME *runtime.Runtime
var COMMAND_CODES = map[int]string {
    0: "start game",
    1: "update tick",
    2: "set queue" }

func ProcessOrders() {
    RUNTIME = runtime.New(GAME)

    for {
        order := <-ORDER_STREAM

        if COMMAND_CODES[order.GetCommandCode()] == "update tick" {
            UpdateTick(order)
        } else if COMMAND_CODES[order.GetCommandCode()] == "set queue" {
            entity := GAME.GetEntityByPlayer(order.GetPlayer())
            entity.SetActionQueue(order.GetActions())
            RUNTIME.Enqueue(order)
        }

        DumpOrder(order)
    }
}

/* Prints the given order */
func DumpOrder(order interfaces.Order) {
    entity := GAME.GetEntityByPlayer(order.GetPlayer())

    log.Printf("Parsed Order: command:'%v' tick:%v actions:%v player:%v",
        COMMAND_CODES[order.GetCommandCode()], order.GetTickNumber(),
        entity.GetStringActionQueue(),
        order.GetPlayer().GetPlayerId())
}

/* Updates players' ticks and send out gamestate when everyone's ready */
func UpdateTick(order interfaces.Order) {
    p := order.GetPlayer()
    p.SetLastTick(order.GetTickNumber())
    p.SetLastTickTime(time.Now())

}

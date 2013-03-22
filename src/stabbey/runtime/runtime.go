package runtime

import (
    "log"
    "time"

    "stabbey/interfaces"
)

var TURN_DELAY, _ = time.ParseDuration("500ms")
var GAME interfaces.Game
var COMMAND_CODES = map[int]string {
    0: "start game",
    1: "update tick",
    2: "set queue" }

type Runtime struct {
    /* Reads from this to handle any orders */
    orderStream chan interfaces.Order
}

func New(game interfaces.Game) *Runtime {
    r := &Runtime{}
    log.Println("CREATING NEW GAME RUNTIME")

    r.orderStream = make(chan interfaces.Order)

    GAME = game
    initLevel(0)

    go r.processOrders()
    go r.scheduleActions()

    return r
}

/* Goroutine to process incoming orders from players */
func (r *Runtime) processOrders() {
    for {
        order := <-r.orderStream

        if COMMAND_CODES[order.GetCommandCode()] == "update tick" {
            updateTick(order)
        } else if COMMAND_CODES[order.GetCommandCode()] == "set queue" {
            entity := GAME.GetEntityByPlayer(order.GetPlayer())
            entity.SetActionQueue(order.GetActions())
            printOrder(order)
        }
    }
}

/* Dumps the given order to the console */
func printOrder(order interfaces.Order) {
    entity := GAME.GetEntityByPlayer(order.GetPlayer())
    log.Printf("Parsed Order: command:'%v' tick:%v actions:%v player:%v",
        COMMAND_CODES[order.GetCommandCode()], order.GetTickNumber(),
        entity.GetStringActionQueue(),
        order.GetPlayer().GetPlayerId())
}

/* Allows for player orders to be added */
func (r *Runtime) AddOrder(order interfaces.Order) {
    r.orderStream <- order
}

/* Goroutine to run queued actions & send updates to players */
func (r *Runtime) scheduleActions() {
    /* The first time all players are ready, we just send the initial state */
    for {
        if allPlayersReady() {
            log.Println("Players ready; sent initial game state")
            GAME.SetLastTick(GAME.GetLastTick() + 1)
            broadcastGamestate()
            break
        }
        time.Sleep(TURN_DELAY)
    }

    /* Thereafter, the players are only ready if they've a queue */
    for {
        if allPlayersReady() == false {
            time.Sleep(TURN_DELAY)
            continue
        }

        for i := 0; i < len(GAME.GetPlayers()); i++ {
            player := GAME.GetPlayer(i)
            entity := GAME.GetEntityByPlayer(player)

            if action := entity.PopAction(); action != nil {
                act(entity, action)
            } else {
                log.Fatalf("Player %v didn't have move ready!", i)
            }

            GAME.SetLastTick(GAME.GetLastTick() + 1)
            broadcastGamestate()
            time.Sleep(TURN_DELAY)
        }

        worldTick()
    }
}


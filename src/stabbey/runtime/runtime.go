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

            if (len(order.GetActions()) <= 0) {
                log.Printf("Ignoring %v's set queue: no actions set",
                    order.GetPlayer().GetName())
                continue
            }

            entity.SetActionQueue(order.GetActions())
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
            GAME.Run()
            GAME.SetLastTick(GAME.GetLastTick() + 1)
            broadcastGamestate()
            break
        }
        time.Sleep(TURN_DELAY)
    }

    /* Thereafter, the players are only ready if they've a queue */
    for {
        /* Ensure all players are ready */
        if allPlayersReady() == false {
            time.Sleep(TURN_DELAY)
            continue
        }

        /* Do all the player actions */
        for i := 0; i < len(GAME.GetPlayers()); i++ {
            player := GAME.GetPlayer(i)
            entity := GAME.GetEntityByPlayer(player)

            if action := entity.PopAction(); action != nil {
                act(entity, action)
            } else {
                log.Printf("Player %v didn't have move ready " +
                    "(client set queue error)!" +
                    " This'll count as 1 do nothing.", i)
            }

            GAME.SetLastTick(GAME.GetLastTick() + 1)
            broadcastGamestate()
            time.Sleep(TURN_DELAY)
        }

        /* Do all the other turns (mostly monster moves) */
        for _, entity := range GAME.GetEntities() {
            if bId, _, _ := entity.GetPosition(); bId == GAME.GetCurrentBoard() {
                if entity.RunTurn(GAME.GetLastTick()) {
                    GAME.SetLastTick(GAME.GetLastTick() + 1)
                    broadcastGamestate()
                    time.Sleep(TURN_DELAY)
                }
            }
        }
    }
}

/* Checks whether all players are ready (done queueing) */
func allPlayersReady() bool {
    players := GAME.GetPlayers()

    /* Has to be at least one ready player */
    if len(players) == 0 {
        return false
    }

    for _, player := range players {
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

    for _, spectator := range GAME.GetSpectators() {
        spectator.SendMessage(GAME.Json(nil))
    }
}

/* Generic action handler for any entity */
func act(entity interfaces.Entity, action interfaces.Action) {
    if err := action.Act(entity, GAME); err != nil {
        log.Printf("%v: %v", entity.GetName(), err)
    }
}

/* Updates players' ticks and send out gamestate when everyone's ready */
func updateTick(order interfaces.Order) {
    p := order.GetPlayer()
    p.SetLastTick(order.GetTickNumber())
    p.SetLastTickTime(time.Now())
}

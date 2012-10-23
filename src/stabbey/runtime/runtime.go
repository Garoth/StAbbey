package runtime

import (
    "log"
    "time"

    "stabbey/interfaces"
)

const DEFAULT_COOLDOWN int64 = 1000
const COOLDOWN_DECREMENT int64 = 50
const MIN_COOLDOWN int64 = 200
const TURN_DELAY string = "100ms"

/* I'm a bit worried about race conditions around GAME in the future */
var GAME interfaces.Game
var COMMAND_CODES = map[int]string {
    0: "start game",
    1: "update tick",
    2: "set queue" }

type timer struct {
    LastCooldown, TimeRemaining int64
}

type Runtime struct {
    /* Reads from this to handle any orders */
    orderStream chan interfaces.Order
    /* Reads from this to insert a new entity queue into the ordering */
    queueOrders chan interfaces.Order
    /* Collection of entities their last cooldown */
    watchedEntities map[interfaces.Entity] *timer
}

func New(game interfaces.Game) *Runtime {
    r := &Runtime{}
    log.Println("CREATING NEW GAME RUNTIME")

    r.orderStream = make(chan interfaces.Order)
    r.queueOrders = make(chan interfaces.Order)
    r.watchedEntities = make(map[interfaces.Entity] *timer, 100)

    GAME = game
    initLevel(0)

    go r.processOrders()
    go r.acceptQueues()
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
            r.AddMoveQueue(order)
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

/* Interfaces for the outside world to insert their orders */
func (r *Runtime) AddOrder(order interfaces.Order) {
    r.orderStream <- order
}

/* Goroutine to accept new move queues for the scheduler */
func (r *Runtime) acceptQueues() {
    for {
        /* We only expect 'set queue' orders to come in */
        order := <-r.queueOrders

        entity := GAME.GetEntityByPlayer(order.GetPlayer())
        r.watchedEntities[entity] = &timer{DEFAULT_COOLDOWN, DEFAULT_COOLDOWN}
    }
}

/* Interface for the outside world to add queues to the scheduler */
func (r *Runtime) AddMoveQueue(order interfaces.Order) {
    r.queueOrders <- order
}

/* Goroutine to run queued actions & send updates to players */
func (r *Runtime) scheduleActions() {
    for {
        turnDelay, _ := time.ParseDuration(TURN_DELAY)
        turnDelayMillis := turnDelay.Nanoseconds() / 1000000
        time.Sleep(turnDelay)

        if allPlayersReady() == false {
            continue
        }

        for entity, timings := range r.watchedEntities {
            timings.TimeRemaining -= turnDelayMillis

            if timings.TimeRemaining <= 0 {
                if timings.LastCooldown > MIN_COOLDOWN {
                    timings.LastCooldown -= turnDelayMillis
                }
                timings.TimeRemaining = timings.LastCooldown

                if action := entity.PopAction(); action != nil {
                    act(entity, action)
                } else {
                    delete(r.watchedEntities, entity)
                }
            }
        }

        worldTick()

        GAME.SetLastTick(GAME.GetLastTick() + 1)
        log.Printf("All players are ready, sending tick %v", GAME.GetLastTick())
        broadcastGamestate()
    }
}

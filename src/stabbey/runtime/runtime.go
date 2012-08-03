package runtime

import (
    "log"
    "time"

    "stabbey/interfaces"
)

const DEFAULT_COOLDOWN int = 500
const COOLDOWN_DECREMENT int = 50
const MIN_COOLDOWN int = 200
const TURN_DELAY string = "1s"

/* I'm a bit worried about race conditions around GAME in the future */
var GAME interfaces.Game
var COMMAND_CODES = map[int]string {
    0: "start game",
    1: "update tick",
    2: "set queue" }

type Runtime struct {
    /* Reads from this to handle any orders */
    orderStream chan interfaces.Order
    /* Reads from this to insert a new entity queue into the ordering */
    queueOrders chan interfaces.Order
    /* Collection of entities their last cooldown */
    watchedEntities map[interfaces.Entity]int
}

func New(game interfaces.Game) *Runtime {
    r := &Runtime{}
    r.orderStream = make(chan interfaces.Order)
    r.queueOrders = make(chan interfaces.Order)
    r.watchedEntities = make(map[interfaces.Entity]int, 100)
    GAME = game
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
        }

        entity := GAME.GetEntityByPlayer(order.GetPlayer())

        log.Printf("Parsed Order: command:'%v' tick:%v actions:%v player:%v",
            COMMAND_CODES[order.GetCommandCode()], order.GetTickNumber(),
            entity.GetStringActionQueue(),
            order.GetPlayer().GetPlayerId())
    }
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
        r.watchedEntities[entity] = DEFAULT_COOLDOWN
    }
}

/* Interface for the outside world to add queues to the scheduler */
func (r *Runtime) AddMoveQueue(order interfaces.Order) {
    r.queueOrders <- order
}

/* Goroutine to run queued actions & send updates to players */
func (r *Runtime) scheduleActions() {
    for {
        if duration, e := time.ParseDuration(TURN_DELAY); e == nil {
            time.Sleep(duration)
        } else {
            log.Fatalf("Error parsing duration, %v", e)
        }

        allReady := true
        for _, player := range GAME.GetPlayers() {
            /* Ready players are 1 ahead of the game's tick */
            if player.GetLastTick() <= GAME.GetLastTick() {
                allReady = false
            }
        }
        if allReady == false {
            continue
        }

        r.executeNext()

        log.Printf("All players are ready, sending next tick")
        GAME.SetLastTick(GAME.GetLastTick() + 1)
        broadcastGamestate()
    }

}

func (r *Runtime) executeNext() {
    /* Run the entity that has the lowest current cooldown */
    highestCooldown := 0
    var nextEntity interfaces.Entity = nil
    for entity, cooldown := range r.watchedEntities {
        if cooldown > highestCooldown {
            highestCooldown = cooldown
            nextEntity = entity
        }
    }

    if nextEntity != nil {
        if action := nextEntity.PopAction(); action != nil {
            act(nextEntity, action)
            r.reduceCooldown(nextEntity)
        } else {
            delete(r.watchedEntities, nextEntity)
            /* Lets try again, no actions left on this one */
            r.executeNext()
        }
    } else {
        log.Printf("Not watching any entities")
    }
}

func (r *Runtime) reduceCooldown(entity interfaces.Entity) {
    cooldown := r.watchedEntities[entity]
    if cooldown > MIN_COOLDOWN {
        r.watchedEntities[entity] = cooldown - COOLDOWN_DECREMENT
    }
}


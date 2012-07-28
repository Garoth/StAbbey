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

type Runtime struct {
    /* Reads from this to insert a new entity queue into the ordering */
    queueOrders chan interfaces.Order
    /* Collection of entities their last cooldown */
    watchedEntities map[interfaces.Entity]int
    /* The game object */
    game interfaces.Game
}

func New(game interfaces.Game) *Runtime {
    r := &Runtime{}
    r.queueOrders = make(chan interfaces.Order)
    r.watchedEntities = make(map[interfaces.Entity]int, 100)
    r.game = game
    go r.run()
    return r
}

/* TODO any race conditions between the adding and manifesting threads? */
func (r *Runtime) run() {
    /* Read new orders in */
    go func() {
        for {
            /* We only expect 'set queue' orders to come in */
            order := <-r.queueOrders

            entity := r.game.GetEntityByPlayer(order.GetPlayer())
            r.watchedEntities[entity] = DEFAULT_COOLDOWN
        }
    }()

    /* Occasionally manifest actions */
    go func() {
        for {
            if duration, e := time.ParseDuration(TURN_DELAY); e == nil {
                time.Sleep(duration)
            } else {
                log.Fatalf("Error parsing duration, %v", e)
            }

            allReady := true
            for _, player := range r.game.GetPlayers() {
                /* Ready players are 1 ahead of the game's tick */
                if player.GetLastTick() <= r.game.GetLastTick() {
                    allReady = false
                }
            }
            if allReady == false {
                continue
            }

            r.executeNext()

            log.Printf("All players are ready, sending next tick")
            r.game.SetLastTick(r.game.GetLastTick() + 1)
            r.broadcastGamestate()
        }
    }()
}

func (r *Runtime) Enqueue(order interfaces.Order) {
    r.queueOrders <- order
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
            Act(nextEntity, action)
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

/* Sends the gamestate to everyone */
func (r *Runtime) broadcastGamestate() {
    for _, player := range r.game.GetPlayers() {
        player.SendMessage(r.game.Json(player))
    }
}

func Act(entity interfaces.Entity, action interfaces.Action) {
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

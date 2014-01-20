package interfaces

import (
    "time"
)

type Player interface {
    Spectator
    Entity
    /* Unique Game ID Getters / Setters */
    GetPlayerId() int
    SetPlayerId(id int)
    /* Available actions getters / setters */
    GetAvailableActions() []Action
    AddAvailableAction(Action)
    /* Last Sent Tick Getters / Setters */
    GetLastTick() int
    SetLastTick(tickNum int)
    /* Last Sent Tick Time Getters / Setters */
    GetLastTickTime() time.Time
    SetLastTickTime(t time.Time)
}

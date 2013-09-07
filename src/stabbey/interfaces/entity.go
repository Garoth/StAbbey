package interfaces

import (
)

const ENTITY_TYPE_PLAYER     = "player"
const ENTITY_TYPE_MONSTER    = "monster"

/* A monster, player, or some special thing of that sort */
type Entity interface {
    /* Entity Ids must be unique */
    SetEntityId(id int)
    GetEntityId() int
    SetPosition(boardid, x, y int)
    GetPosition() (boardid, x, y int)
    SetName(name string)
    GetName() string
    SetType(t string)
    GetType() string
    SetMaxArdour(ardour int)
    GetMaxArdour() int
    ChangeArdour(difference int) int
    SetArdour(ardour int)
    GetArdour() int
    IsDead() bool
    /* Queue Manipulation Getters / Setters */
    GetActionQueue() []Action
    GetStringActionQueue() []string
    SetActionQueue([]Action)
    PopAction() Action
}

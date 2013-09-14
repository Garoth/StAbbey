package interfaces

import (
)

const (
    ENTITY_TYPE_PLAYER            = "player"
    ENTITY_TYPE_MONSTER           = "monster"
    ENTITY_TYPE_TRIGGER           = "loot"

    TRIGGER_TYPE_ABILITY_PUSH     = "pu"
)

/* A monster, player, or some special thing of that sort */
type Entity interface {
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
    IsTangible() bool
    SetTangible(tangible bool)
    IsDead() bool
    Die()
    Trodden(by Entity)
    GetActionQueue() []Action
    GetStringActionQueue() []string
    SetActionQueue([]Action)
    PopAction() Action
    WorldTick(tick int)
}

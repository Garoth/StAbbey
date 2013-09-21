package interfaces

import (
)

const (
    ENTITY_TYPE_UNKNOWN                  = "unknown type"
    ENTITY_TYPE_PLAYER                   = "player"
    ENTITY_TYPE_MONSTER                  = "monster"
    ENTITY_TYPE_TRIGGER                  = "trigger"
    ENTITY_TYPE_INERT                    = "inert"

    ENTITY_SUBTYPE_UNKNOWN               = "unknown subtype"
    ENTITY_MONSTER_SUBTYPE_GARGOYLE      = "gargoyle"
    ENTITY_MONSTER_SUBTYPE_CHEST         = "chest"
    ENTITY_TRIGGER_SUBTYPE_ABILITY_LOOT  = "ability loot"
    ENTITY_TRIGGER_SUBTYPE_TELEPORT_TRAP = "teleport trap"
    ENTITY_TRIGGER_SUBTYPE_CALTROP_TRAP  = "caltrop trap"
    ENTITY_TRIGGER_SUBTYPE_STAIRS_UP     = "stairs up"
    ENTITY_INERT_SUBTYPE_TRAP            = "sprung trap"
    ENTITY_INERT_SUBTYPE_TREE            = "tree"
    ENTITY_INERT_SUBTYPE_STATUE          = "inert statue"

    TRIGGER_TYPE_ABILITY_PUSH            = "pu"
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
    SetSubtype(t string)
    GetSubtype() string
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

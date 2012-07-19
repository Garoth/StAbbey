package interfaces

import (
)

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
}

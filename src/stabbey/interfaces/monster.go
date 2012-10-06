package interfaces

import (
)

type Monster interface {
    Entity
    /* Unique Game ID Getters / Setters */
    GetMonsterId() int
    SetMonsterId(id int)
}

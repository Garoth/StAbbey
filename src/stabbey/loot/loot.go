package loot

import (
    "stabbey/entity"
)

type Loot struct {
    *entity.Entity
}

func New() *Loot {
    me := &Loot{}
    me.Entity = entity.New(entity.UIDG.NextUid())
    me.Die()
}

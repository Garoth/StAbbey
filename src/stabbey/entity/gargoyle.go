package entity

import (
    "strconv"
    "stabbey/interfaces"
)

func NewGargoyle(g interfaces.Game) interfaces.Entity {
    me := newBasicMonster(g)
    me.SetSubtype(interfaces.ENTITY_MONSTER_SUBTYPE_GARGOYLE)
    me.SetName("Gargoyle " + strconv.Itoa(me.GetEntityId()))
    me.SetMaxArdour(30)
    me.SetArdour(30)

    return me
}

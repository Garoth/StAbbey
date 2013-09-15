package entity

import (
    "strconv"
    "stabbey/interfaces"
)

func NewInertStatue(g interfaces.Game) *Entity {
    me := newBasicInert(g)
    me.SetName("Inert Statue " + strconv.Itoa(me.GetEntityId()))
    me.SetMaxArdour(100)
    me.SetArdour(100)
    me.SetSubtype(interfaces.ENTITY_INERT_SUBTYPE_STATUE)
    return me
}

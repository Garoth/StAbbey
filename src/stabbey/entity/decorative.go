package entity

import (
    "strconv"
    "stabbey/interfaces"
)

func NewTree(g interfaces.Game) *Entity {
    me := newBasicInert(g)
    me.SetName("Tree " + strconv.Itoa(me.GetEntityId()))
    me.SetMaxArdour(80)
    me.SetArdour(80)
    me.SetSubtype(interfaces.ENTITY_INERT_SUBTYPE_TREE)
    return me
}

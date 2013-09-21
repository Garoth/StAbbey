package entity

import (
    "log"
    "strconv"
    "stabbey/interfaces"
)

func NewStairsUp(g interfaces.Game) *Entity {
    me := newBasicTrigger(g)
    me.SetName("Stairs Up " + strconv.Itoa(me.GetEntityId()))
    me.SetMaxArdour(10000)
    me.SetArdour(10000)
    me.SetSubtype(interfaces.ENTITY_TRIGGER_SUBTYPE_STAIRS_UP)

    me.TroddenFunction = func(by interfaces.Entity) {
        log.Println(by.GetName(), "reached stairs. Loading next board.")
        me.Game.NextBoard()
    }

    return me
}

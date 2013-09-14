package entity

import (
    "strconv"
    "stabbey/interfaces"
)

func NewGargoyle(g interfaces.Game) interfaces.Entity {
    me := newBasicMonster(g)
    me.SetName("Gargoyle " + strconv.Itoa(me.GetEntityId()))
    me.SetMaxArdour(30)
    me.SetArdour(30)

    me.TickFunction = func(tick int) {
    }

    return me
}

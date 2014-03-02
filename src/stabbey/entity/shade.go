package entity

import (
    "strconv"
    "stabbey/interfaces"
)

func NewShade(g interfaces.Game) interfaces.Entity {
    me := newBasicMonster(g)
    me.SetSubtype(interfaces.ENTITY_MONSTER_SUBTYPE_SHADE)
    me.SetName("Shade " + strconv.Itoa(me.GetEntityId()))
    me.SetMaxArdour(15)
    me.SetArdour(15)

    me.TurnFunction = func(tick int) bool {
        didStuff := false


        return didStuff
    }

    return me
}

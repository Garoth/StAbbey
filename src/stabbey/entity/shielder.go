package entity

import (
    "strconv"
    "stabbey/interfaces"
    "stabbey/order"
)

func NewShielder(g interfaces.Game) interfaces.Entity {
    me := newBasicMonster(g)
    me.SetSubtype(interfaces.ENTITY_MONSTER_SUBTYPE_SHIELDER)
    me.SetName("Shielder " + strconv.Itoa(me.GetEntityId()))
    me.SetMaxArdour(50)
    me.SetArdour(50)

    push := make(map[string]interfaces.Action)
    push["left"] = order.NewAction("pl")
    push["right"] = order.NewAction("pr")
    push["up"] = order.NewAction("pu")
    push["down"] = order.NewAction("pd")

    me.TurnFunction = func(tick int) bool {
        foundAny := false

        for _, action := range push {
            if err := action.Act(me, me.Game); err == nil {
                foundAny = true
            }
        }

        return foundAny
    }

    return me
}

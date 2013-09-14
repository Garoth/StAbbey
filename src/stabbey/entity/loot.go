package entity

import (
    "log"
    "stabbey/interfaces"
    "stabbey/order"
)

func NewAbilityTrigger(g interfaces.Game, abilityCode string) *Entity {
    me := newBasicTrigger(g)
    me.SetSubtype(interfaces.ENTITY_TRIGGER_SUBTYPE_ABILITY_LOOT)
    me.SetName("Loot for ability " + abilityCode)

    me.TroddenFunction = func(by interfaces.Entity) {
        log.Println(me.GetName() + " trodden on by", by.GetName())

        if by.GetType() == interfaces.ENTITY_TYPE_PLAYER {
            player := me.Game.GetPlayerByEntity(by)
            player.AddAvailableAction(order.NewAction(abilityCode))
            me.SetArdour(0)
        }
    }

    return me
}

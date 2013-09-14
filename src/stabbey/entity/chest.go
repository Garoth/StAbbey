package entity

import (
    "log"
    "strconv"
    "stabbey/interfaces"
)

func NewChest(g interfaces.Game) interfaces.Entity {
    me := newBasicMonster(g)
    me.SetSubtype(interfaces.ENTITY_MONSTER_SUBTYPE_CHEST)
    me.SetName("Chest " + strconv.Itoa(me.GetEntityId()))
    me.SetMaxArdour(10)
    me.SetArdour(10)

    me.TickFunction = func(tick int) {
    }

    me.DeathFunction = func() {
        log.Println(me.GetName(), "drops loot")
        loot := NewAbilityTrigger(me.Game, interfaces.TRIGGER_TYPE_ABILITY_PUSH)
        loot.SetPosition(me.GetPosition())
        me.Game.AddEntity(loot)
    }

    return me
}

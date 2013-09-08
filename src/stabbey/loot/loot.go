package loot

import (
    "log"
    "stabbey/entity"
    "stabbey/interfaces"
    "stabbey/order"
)

type Loot struct {
    *entity.Entity
    LootType string
    GameFunctions interfaces.LootModifyGame
}

func New() *Loot {
    me := &Loot{}

    /* Loot Stuff */
    me.LootType = interfaces.LOOT_TYPE_ABILITY_PUSH

    /* Entity Stuff */
    me.Entity = entity.New(entity.UIDG.NextUid())
    me.SetName("Loot for " + me.LootType)
    me.SetType(interfaces.ENTITY_TYPE_LOOT)
    me.SetTangible(false)
    me.TroddenFunction = func(by interfaces.Entity) {
        log.Println("Loot for", me.LootType, "trodden on by", by.GetName())

        if by.GetType() == interfaces.ENTITY_TYPE_PLAYER {
            player := me.GameFunctions.GetPlayerByEntity(by)
            player.AddAvailableAction(order.NewAction(me.LootType))
            me.SetArdour(0)
        }
    }

    return me;
}

func (me *Loot) SetGameFunctions(gameFns interfaces.LootModifyGame) {
    me.GameFunctions = gameFns
}

package entity

import (
    "log"
    "strconv"
    "stabbey/interfaces"
)

func NewTeleportTrap(g interfaces.Game, x, y int) *Entity {
    me := newBasicTrigger(g)
    me.SetSubtype(interfaces.ENTITY_TRIGGER_SUBTYPE_TELEPORT_TRAP)
    me.SetName("Teleport trap " + strconv.Itoa(me.GetEntityId()) +  " to " +
        strconv.Itoa(x) + ", " + strconv.Itoa(y))

    me.TroddenFunction = func(by interfaces.Entity) {
        /* I've already triggered */
        if me.IsDead() {
            return
        }

        if me.Game.CanMoveToSpace(g.GetCurrentBoard(), x, y) {
            log.Println("Teleporting", by.GetName(), "to", x, y)
            by.SetPosition(me.Game.GetCurrentBoard(), x, y)
        } else {
            log.Println(me.GetName() + " failed, since destination is blocked")
        }
        me.Die()

        sprungTrap := NewSprungTrap(g)
        sprungTrap.SetPosition(me.GetPosition())
        me.Game.AddEntity(sprungTrap)
    }

    return me
}

func NewCaltropTrap(g interfaces.Game) *Entity {
    damage := 20
    me := newBasicTrigger(g)
    me.SetSubtype(interfaces.ENTITY_TRIGGER_SUBTYPE_CALTROP_TRAP)
    me.SetName("Caltrop trap " + strconv.Itoa(me.GetEntityId()))

    me.TroddenFunction = func(by interfaces.Entity) {
        if by.IsTangible() {
            log.Println(me.GetName())
            by.ChangeArdour(-damage)
        }
    }

    return me
}

func NewSprungTrap(g interfaces.Game) *Entity {
    me := newBasicInert(g)
    me.SetSubtype(interfaces.ENTITY_INERT_SUBTYPE_TRAP)
    me.SetTangible(false)
    me.SetName("Sprung trap")
    return me
}

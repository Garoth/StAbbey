package entity

import (
    "log"
    "strconv"
    "stabbey/interfaces"
)

func NewTeleportTrap(g interfaces.Game, x, y int) *Entity {
    me := newBasicTrigger(g)
    me.SetSubtype(interfaces.ENTITY_TRIGGER_SUBTYPE_TELEPORT_TRAP)
    me.SetName("Teleport trap to " + strconv.Itoa(x) + ", " + strconv.Itoa(y))

    me.TroddenFunction = func(by interfaces.Entity) {
        log.Println(me.GetName() + " trodden on by", by.GetName())

        /* I've already triggered */
        if me.IsDead() {
            return
        }

        if me.Game.CanMoveToSpace(x, y) {
            log.Println("Teleporting", by.GetName(), "to", x, y)
            by.SetPosition(me.Game.GetCurrentBoard(), x, y)
        } else {
            log.Println(me.GetName() + " failed, since destination is blocked")
        }
        me.Die()

        _, myX, myY := me.GetPosition()
        sprungTrap := NewSprungTrap(g, myX, myY)
        me.Game.AddEntity(sprungTrap)
    }

    return me
}

func NewSprungTrap(g interfaces.Game, x, y int) *Entity {
    me := newBasicInert(g)
    me.SetSubtype(interfaces.ENTITY_INERT_SUBTYPE_TRAP)
    me.SetTangible(false)
    me.SetPosition(me.Game.GetCurrentBoard(), x, y)
    me.SetName("Sprung trap at " + strconv.Itoa(x) + ", " + strconv.Itoa(y))
    return me
}
